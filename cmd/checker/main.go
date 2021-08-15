package main

import (
	"log"
	"os"
	"time"

	"nodeChecker/internal/alerts"
	"nodeChecker/internal/checker"
)

func main() {
	key := os.Args[1]
	username := os.Args[2]
	alertManager := alerts.NewAlertManager(username, key)
	nodeChecker, err := checker.NewNodePortChecker()
	if err != nil {
		log.Println(err)
		alertManager.SendAlert(err.Error(), "P1")
		return
	}
	trdChecker, err := checker.NewTRDChecker()
	if err != nil {
		log.Println(err)
		alertManager.SendAlert(err.Error(), "P1")
		return
	}
	checkers := []checker.Checker{nodeChecker, trdChecker}

	// chckers loop
	for {
		for _, ch := range checkers {
			if err = ch.AssertRunning(); err != nil {
				log.Println(err)
				if err = alertManager.SendAlert(err.Error(), "P1"); err != nil {
					log.Fatalln(err)
				}
			}
			time.Sleep(time.Minute * time.Duration(10))
		}
	}
}

//func bootstrapAlert(err error, key string, username string) (error, bool) {
//	if err != nil {
//		log.Println(err)
//		if err = alertManager.SendAlert(key, username, err.Error(), "P1"); err != nil {
//			log.Fatalln(err)
//		}
//		return nil, true
//	}
//	return err, false
//}
