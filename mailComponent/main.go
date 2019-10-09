package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	configs "../configs"
	logging "../logging"

	persistence "../persistence"
)

var appMode = "DEVELOPMENT"

func init() {
	configSetUpArgs := &configs.ConfigArgs{
		ConfigFilePath: []string{"./configs", "$HOME/configs"},
		ConfigfileName: "appconfigs",
	}
	configs.SetUp(configSetUpArgs)
}

func main() {
	// App Start, Printing banner
	printBanner()

	//created channel to collect signals from os
	osSignals := make(chan os.Signal, 1)
	defer close(osSignals)

	//created channel to collect value, getting value from this channel means terminate app
	terminateApp := make(chan bool, 1)
	defer close(terminateApp)

	// TO register channel so that os signal can be passed to this channel
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-osSignals
		fmt.Printf("got the sigal %v", sig)
		terminateApp <- true
	}()

	//conf := configs.Get("LOGGING")
	//fmt.Printf("%v", conf)
	logging.Setup(configs.Get("LOGGING"))
	persistence.SetUp(configs.Get("DB"))

	// Doing initial setup For Logging, DB and NATS

	go printTime()

	// blocking to terminate the application once it collect any value it would be close
	<-terminateApp
	fmt.Println("terminating the application")
	fmt.Println("Application termiated")

}

func printBanner() {
	contentBytes, error := ioutil.ReadFile("./configs/banner.txt")
	if error == nil {
		fmt.Println(string(contentBytes))
		fmt.Printf("AppMode: %v", appMode)
		fmt.Println(" ")
	} else {
		fmt.Println(fmt.Errorf("some error occurred while reading file %v", error))
	}
}
