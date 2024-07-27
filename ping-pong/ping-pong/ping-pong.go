package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	_ "github.com/lib/pq"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var mu sync.Mutex

const (
	pingLogFilePath = "/pinglogs/pinglog.txt"
	host            = "postgres-svc.applications.svc.cluster.local"
	user            = "postgres"
	dbname          = "db"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCounter(w http.ResponseWriter, db *sql.DB) int {
	var counter int
	err := db.QueryRow("SELECT counter FROM db.ping_logs ORDER BY id DESC LIMIT 1 FOR UPDATE").Scan(&counter)

	if err == nil {
		fmt.Fprintf(w, fmt.Sprintf("pong %s", fmt.Sprintf("%v", counter)))
		return counter
	} else {
		panic(err)
	}
}

func getPostgresPassword() string {
	ctx := context.Background()
	config, err := rest.InClusterConfig()
	if err != nil {
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		configOverrides := &clientcmd.ConfigOverrides{}
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
		config, err = kubeConfig.ClientConfig()
		check(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	check(err)

	secret, err := clientset.CoreV1().Secrets("applications").Get(ctx, "postgres-key", metav1.GetOptions{})
	check(err)

	var password = string(secret.Data["POSTGRES_PASS"])

	return password
}

func main() {
	password := getPostgresPassword()
	psqlInfo := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Print("Postgres connection successful")
	}

	db.Exec("INSERT INTO db.ping_logs (counter) VALUES ($1)", 0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		counter := getCounter(w, db)
		counter++
		db.Exec("DELETE FROM db.ping_logs WHERE id = (SELECT id FROM db.ping_logs ORDER BY id DESC LIMIT 1)")
		db.Exec("INSERT INTO db.ping_logs (counter) VALUES ($1)", counter)
	})
	http.HandleFunc("/no-increment", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		getCounter(w, db)
	})

	port := os.Getenv("PORT")
	fmt.Printf("Server started in port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
