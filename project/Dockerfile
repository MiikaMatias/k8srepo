FROM golang:latest

RUN mkdir /go/assets

COPY server.go /go/
COPY index.html /go/
RUN chmod 777 -R /go

WORKDIR /go

CMD ["go", "run", "server.go"]