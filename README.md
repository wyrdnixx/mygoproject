# Go Docker Workspace template

~/.profile   (Anschließend # source .profile)
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin

---------------------------------------------------


ulewu@ubu:~/go/src/github.com/wyrdnixx/go-docker-template$ tree
.
├── api
├── cmd
│   └── api
│       ├── Dockerfile
│       ├── main.go
│       └── modules
│           └── modules.go
├── docker-compose.yml
├── Makefile
└── README.md

-----------------------------------------

Pfad / Paketname anpassen:

-> ~/go/src/github.com/wyrdnixx/<go-docker-template>/cmd/api/Dockerfile

WORKDIR /go/src/github.com/wyrdnixx/<go-docker-template>


-> ~/go/src/github.com/wyrdnixx/<go-docker-template>/docker-compose.yml
        volumes:
            - ./:/go/src/github.com/wyrdnixx/<go-docker-template>


-> ~/go/src/github.com/wyrdnixx/<go-docker-template>/cmd/api/main.go
	"github.com/wyrdnixx/<go-docker-template>/cmd/api/modules"
-----------------------------------------


Docker benutzen:

# cd ~/go/src/github.com/wyrdnixx/<go-docker-template>
# make up
# make logs
# make clean