package main

import (
	"log"

	"github.com/ffajarpratama/gommerce-api/cmd/app"
)

func main() {
	if err := app.Exec(); err != nil {
		log.Fatal("[app-failed]", err)
	}
}
