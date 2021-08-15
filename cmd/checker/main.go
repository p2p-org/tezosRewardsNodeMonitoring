package main

import (
	"log"
	"os"
	"time"

	"nodeChecker/internal/alerts"
	"nodeChecker/internal/checker"
	"nodeChecker/internal/fswatch"
)

func main() {
	key := os.Args[1]
	username := os.Args[2]
	folder := os.Args[3]
	alertManager := alerts.NewAlertManager(username, key)
	reportsWatcher := fswatch.NewReportsWatcher(folder, alertManager)
	go reportsWatcher.Watch()
	nodeChecker, err := checker.NewNodePortChecker()
	if err != nil {
		log.Println(err)
		alertManager.SendAlert("bootsrap error", err.Error(), "P1")
		return
	}
	trdChecker, err := checker.NewTRDChecker()
	if err != nil {
		log.Println(err)
		alertManager.SendAlert("bootstrap error", err.Error(), "P1")
		return
	}
	checkers := []checker.Checker{nodeChecker, trdChecker}

	// chckers loop
	for {
		for _, ch := range checkers {
			if err = ch.AssertRunning(); err != nil {
				log.Println(err)
				if err = alertManager.SendAlert(ch.GetTitle(), err.Error(), "P1"); err != nil {
					log.Fatalln(err)
				}
			}
			time.Sleep(time.Minute * time.Duration(10))
		}
	}
}
