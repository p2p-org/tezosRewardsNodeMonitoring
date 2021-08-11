package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"

	"nodeChecker/internal/checker"
)

func main() {
	key := os.Args[1]
	username := os.Args[2]
	var config client.Config
	config.ApiKey = key
	alertCli, err := alert.NewClient(&config)
	if err != nil {
		log.Panic(err)
	}
	nodeChecker, err := checker.NewNodePortChecker()
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
				responder := alert.Responder{
					Type:     alert.UserResponder,
					Username: username,
				}
				req := alert.CreateAlertRequest{
					Message:    err.Error(),
					Priority:   alert.P1,
					Responders: []alert.Responder{responder},
				}
				if resp, err := alertCli.Create(context.Background(), &req); err != nil {
					log.Println(resp.Result)
					log.Panic(err)
				}
			}
			time.Sleep(time.Minute * time.Duration(10))
		}
	}
}
