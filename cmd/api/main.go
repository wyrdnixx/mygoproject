package main

import (
	"fmt"
	"log"
	"net/http"

	//gonfig -> config aus json file lesen

	"github.com/wyrdnixx/mygoproject/cmd/api/modules"
)

//var AppConfig = modules.Configuration{}

func main() {

	modules.HelloWorld()

	modules.ReadConfig()

	fmt.Println("main() Info: ", modules.AppConfig.Info)

	modules.CheckDB()

	// Mit dem Export des Typs "AppConfig" kann jetzt auch von aussen zugegriffen werden
	//	modules.AppConfig.Info = "Neue Info"
	// 	fmt.Println("main() Info: ", modules.AppConfig.Info)

	http.HandleFunc("/", modules.MainHandler)
	log.Println("listening on: ", modules.AppConfig.SrvPort)
	log.Fatal(http.ListenAndServe(modules.AppConfig.SrvPort, nil))
}
