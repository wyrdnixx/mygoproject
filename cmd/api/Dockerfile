FROM golang:1.10

WORKDIR /go/src/github.com/wyrdnixx/mygoproject

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "github.com/tkanos/gonfig"]
RUN ["go", "get", "github.com/lib/pq"]


ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"
