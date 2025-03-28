package cmd

import (
	"log"
	"poll_bot/internal"
)

func main() {
	err := internal.Run()
	if err != nil {
		log.Fatal("fail to run")
	}
}
