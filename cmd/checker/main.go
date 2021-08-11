package main

import (
	"fmt"
	"log"
	"os"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	// "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/policy"

	"nodeChecker/internal/checker"
)

func main() {
	nodeChecker, err := checker.NewNodePortChecker()
	key := os.Args[1]
	if err != nil {
		log.Panic(err)
	}
	_, err = policy.NewClient(&client.Config{ApiKey: key})
	if err != nil {
		fmt.Println("error occured while creating policy client")
		return
	}
	trdChecker, err := checker.NewTRDChecker()
	if err != nil {
		log.Panic(err)
	}
	checkers := []checker.Checker{nodeChecker, trdChecker}

	// chckers loop
	for {
		for _, ch := range checkers {
			if err = ch.AssertRunning(); err != nil {
				log.Println(err)
			}
			// time.Sleep(time.Minute)
		}
	}
}
