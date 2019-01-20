package main

import (
	"fmt"
	"log"
	"net/http"

	//gonfig -> config aus json file lesen
	"github.com/tkanos/gonfig"

	"github.com/wyrdnixx/mygoproject/cmd/api/modules"
)

type Configuration struct {
	SrvPort string
}

func main() {

	modules.HelloWorld()

	// Config aus file laden:
	configuration := Configuration{}
	err := gonfig.GetConf("./cmd/api/config.development.json", &configuration)
	if err != nil {
		fmt.Println("ERROR: Config konnte nicht geladen werden.")
	}

	http.HandleFunc("/", modules.MainHandler)
	log.Println("listening on: ", configuration.SrvPort)
	log.Fatal(http.ListenAndServe(configuration.SrvPort, nil))
}
