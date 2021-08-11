package main

import (
	"log"

	"nodeChecker/internal/checker"
)

func main() {
	nodeChecker, err := checker.NewNodePortChecker()
	if err != nil {
		log.Panic(err)
	}
	trdChecker, err := checker.NewTRDChecker()
	if err != nil {
		log.Panic(err)
	}
	checkers := []checker.Checker{nodeChecker, trdChecker}

	// chckers loop
	for _, ch := range checkers {
		if err = ch.AssertRunning(); err != nil {
			log.Println(err)
		}
	}
}
