package main

import (
	"time"

	log "../logging"
)

func printTime() {

	for {
		time.Sleep(100 * time.Millisecond)
		log.Info(time.Now())
	}

}
