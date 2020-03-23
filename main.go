package main

import (
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

	err := ioutil.WriteFile(tmpFileName, nil, 0644)
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for range c {
		err = os.Remove(tmpFileName)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func fileExists() bool {
	info, err := os.Stat(tmpFileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
