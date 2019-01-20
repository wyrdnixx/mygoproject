FROM golang:1.10

WORKDIR /go/src/github.com/wyrdnixx/mygoproject

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "github.com/tkanos/gonfig"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"
