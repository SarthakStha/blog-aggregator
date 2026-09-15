package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	currState, err := initializeState()
	if err != nil {
		log.Fatalf("Unable to initialize the state: %s", err)
	}

	registry := initializeRegistry()

	if len(os.Args) < 2 {
		log.Fatal("No command provided")
	}

	if _, ok := no_arg_commands[os.Args[1]]; !ok && len(os.Args) < 3 {
		log.Fatal("Insufficient number of argument")
	}

	if err := registry.run(currState, command{os.Args[1], os.Args[2:]}); err != nil {
		log.Fatalf("Error while handling the command: %s", err)
	}

	fmt.Println(currState.cfg.DbURL, currState.cfg.CurrentUserName)
}
