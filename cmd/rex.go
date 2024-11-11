package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mrflobow/rex/services"
)

func main() {
	serverNamePtr := flag.String("server", "", "s")
	flag.Parse()

	fmt.Println("Remote Execution Automator")
	configLoader := services.ConfigLoader{}
	config, err := configLoader.LoadConfig("config.yml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server: %v \n", *serverNamePtr)
	fmt.Println("Args: ", flag.Args())

	exec := services.RemoteExecutor{}
	out, err := exec.ExecuteCommand(config, *serverNamePtr, flag.Args())

	if err != nil {
		log.Fatal(err)
	}

	println(string(out.Data))

}
