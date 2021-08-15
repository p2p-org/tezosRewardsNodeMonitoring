package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"nodeChecker/internal/checker"
)

const (
	// lamp createAlert --message "host down" --config lamp.conf
	//--users pavel.sh@p2p.org --priority P1 --teams Tezos-Team
	lamp = "lamp"
	cmd  = "lamp createAlert --message %s --apiKey %s " +
		"--users %s --priority %s"
)

func sendAlert(key, user, message, priority string) (err error) {
	run := fmt.Sprintf(cmd, message, key, user, priority)
	// log.Println(run)
	_, err = exec.Command("bash", "-c", run).Output()
	// log.Println(string(output))
	return err
}

func main() {
	key := os.Args[1]
	username := os.Args[2]
	//var config client.Config
	//config.ApiKey = key
	// alertCli, err := alert.NewClient(&config)
	//if err != nil {
	//	sendAlert(key, username, err.Error(), "P1")
	//	return
	//}
	nodeChecker, err := checker.NewNodePortChecker()
	if err != nil {
		if err = sendAlert(key, username, err.Error(), "P1"); err != nil {
			log.Fatalln(err)
		}
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
				if err = sendAlert(key, username, err.Error(), "P1"); err != nil {
					log.Fatalln(err)
				}
			}
			time.Sleep(time.Minute * time.Duration(10))
		}
	}
}

//
//func sendAlert(username string, err error, alertCli *alert.Client) {
//	//responder := alert.Responder{
//	//	Type: alert.UserResponder,
//	//	Name: username,
//	//}
//	req := alert.CreateAlertRequest{
//		Message:  err.Error(),
//		Priority: alert.P1,
//		User:     username,
//		//Responders: []alert.Responder{responder,
//		//	{
//		//		Type: alert.TeamResponder,
//		//		Name: "Tezos-Team",
//		//	},
//		//},
//		Responders: [] alert.Responder{
//			{Type: alert.UserResponder, Username: username},
//			{Type: alert.TeamResponder, Name: username},
//		},
//	}
//	resp, err := alertCli.Create(context.Background(), &req)
//	if err != nil {
//		log.Panic(err)
//	}
//	log.Println(resp.Result)
//}
