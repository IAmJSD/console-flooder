package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var text = []byte("Hello, world!\n")

func main() {
	everyMs := os.Getenv("EVERY_MS")
	if everyMs == "" {
		everyMs = "2"
	}
	every, err := strconv.Atoi(everyMs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, _ = os.Stdout.Write(text)
		time.Sleep(time.Duration(every) * time.Second)
	}
}
