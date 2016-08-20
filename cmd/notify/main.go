package main

import (
	"log"
	"os"
	"strings"

	"github.com/twstrike/gosx-notifier"
)

func main() {
	notification := strings.Join(os.Args[1:], " ")

	note := gosxnotifier.NewNotification(notification)

	note.Title = "Notify"

	err := note.Push()

	//If necessary, check error
	if err != nil {
		log.Println("Uh oh! Error with Notify")
	}
}
