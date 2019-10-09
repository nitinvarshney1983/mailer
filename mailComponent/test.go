package main

import (
	"time"

	log "github.com/nitinvarshney1983/mailer/logging"
)

func printTime() {

	for {
		time.Sleep(100 * time.Millisecond)
		log.Info(time.Now())

	}

}
