package main

import (
	"fmt"
	"github.com/SarthakStha/blog-aggregator/internal/config"
	"log"
)

func main() {
	userConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if err := (&userConfig).SetUser("test user"); err != nil {
		log.Fatal(err)
	}

	userConfig, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userConfig.DbURL, userConfig.CurrentUserName)
}
