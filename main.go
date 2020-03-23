package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const tmpFileName = "lock"

func main() {
	if fileExists() {
		log.Fatal("Lock file exists")
	}

	fmt.Println("Writing lock file")
	err := ioutil.WriteFile(tmpFileName, nil, 0644)
	if err != nil {
		log.Fatal(err)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Cleaning up lock file...")
		err = os.Remove(tmpFileName)
		if err != nil {
			log.Fatal(err)
		}
		done <- true
	}()

	fmt.Println("Awaiting signal")
	<-done
	fmt.Println("Exit")
}

func fileExists() bool {
	info, err := os.Stat(tmpFileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
