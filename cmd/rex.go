package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mrflobow/rex/services"
)

func main() {
	serverNamePtr := flag.String("s", "", "Server to use")
	grpNamePtr := flag.String("g", "", "Group to use")
	flag.Parse()

	log.Println("Remote Execution Automator")
	configLoader := services.ConfigLoader{}
	config, err := configLoader.LoadConfig("config.yml")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded - OK")

	if *serverNamePtr != "" && *grpNamePtr != "" {
		log.Fatal("Please choose either server (-s) or Group (-g) option")
	}

	exec := services.RemoteExecutor{Config: config}

	if *serverNamePtr != "" {

		log.Printf("Executing on %v", *serverNamePtr)
		out, err := exec.ExecuteCommand(*serverNamePtr, flag.Args())

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out.Data))
	}

	if *grpNamePtr != "" {
		exec.MultiExec(*grpNamePtr, flag.Args())
	}

}
