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
	cmd = "lamp createAlert --message %s --apiKey %s " +
		"--users %s --priority %s"
)

func sendAlert(key, user, message, priority string) (err error) {
	run := fmt.Sprintf(cmd, message, key, user, priority)
	_, err = exec.Command("bash", "-c", run).Output()
	return err
}

func main() {
	key := os.Args[1]
	username := os.Args[2]
	nodeChecker, err := checker.NewNodePortChecker()
	err, done := bootstrapAlert(err, key, username)
	if done {
		return
	}
	trdChecker, err := checker.NewTRDChecker()
	err, done = bootstrapAlert(err, key, username)
	if done {
		return
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

func bootstrapAlert(err error, key string, username string) (error, bool) {
	if err != nil {
		log.Println(err)
		if err = sendAlert(key, username, err.Error(), "P1"); err != nil {
			log.Fatalln(err)
		}
		return nil, true
	}
	return err, false
}
