package main

import (
	"fmt"
	"time"
)

func printTime() {

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(time.Now())
	}

}
