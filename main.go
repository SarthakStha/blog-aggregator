package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	currState, err := initializeState()
	if err != nil {
		log.Fatalf("Unable to initialize the state: %s", err)
	}

	registry := initializeRegistry()
	registry.register("login", handlerLogin)

	if len(os.Args) < 3 {
		log.Fatal("Insufficient number of argument")
	}
	if err := registry.run(currState, command{os.Args[1], os.Args[2:]}); err != nil {
		log.Fatalf("Error while handling the command: %s", err)
	}

	fmt.Println(currState.currConfig.DbURL, currState.currConfig.CurrentUserName)
}
