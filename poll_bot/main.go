package main

import (
	"log"

	"poll_bot/internal"
)

func main() {
	err := internal.Run()
	if err != nil {
		log.Fatal("fail to run", err.Error())
	}
}
