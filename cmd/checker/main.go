package main

import (
	"fmt"
	"log"
	"os"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
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
	policyClient, err := policy.NewClient(&client.Config{ApiKey: key})
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
				flag := false
				//create a request object
				createAlertPolicyRequest := &policy.CreateAlertPolicyRequest{
					MainFields: policy.MainFields{
						Name:              "my alert policy",
						Enabled:           &flag,
						PolicyDescription: "a policy",
					},
					Message:  err.Error(),
					Continue: &flag,
					Alias:    "test",
					Priority: alert.P1,
				}

				//function call to process the request
				createAlertPolicyResult, err := policyClient.CreateAlertPolicy(nil, createAlertPolicyRequest)

				if err != nil {
					fmt.Printf("error: %s\n", err)
				} else {
					fmt.Printf("result: %+v: ", createAlertPolicyResult)
				}
			}
			// time.Sleep(time.Minute)
		}
	}
}
