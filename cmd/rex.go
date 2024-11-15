package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mrflobow/rex/models"
	"github.com/mrflobow/rex/services"
)

func main() {

	var config *models.Config
	var err error

	serverNamePtr := flag.String("s", "", "Server to use")
	grpNamePtr := flag.String("g", "", "Group to use")
	configFilePtr := flag.String("c", "", "Config file to load, default is <USERHOME>/.rex/config.yml")
	flag.Parse()

	log.Println("REX v0.1")

	configLoader := services.ConfigLoader{}

	if *configFilePtr != "" {
		if config, err = configLoader.LoadConfig(*configFilePtr); err != nil {
			log.Fatal(err)
		}
	} else {
		if config, err = configLoader.LoadDefault(); err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded - OK")

	if *serverNamePtr != "" && *grpNamePtr != "" {
		log.Fatal("Please choose either server (-s) or Group (-g) option")
	}

	exec := services.NewRemoteExecutor(config)

	if *serverNamePtr != "" {

		log.Printf("Executing on %v", *serverNamePtr)
		out, err := exec.ExecuteCommand(*serverNamePtr, flag.Args())

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out.Data))
	}

	if *grpNamePtr != "" {
		out, err := exec.MultiExec(*grpNamePtr, flag.Args())

		if err != nil {
			log.Fatal(err)
		}

		for _, result := range *out {
			fmt.Printf("Server: %v\nOutput:\n%v\n", result.Server, string(result.Data))
		}

	}

}
