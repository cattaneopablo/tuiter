package main

import (
	"log"
	"github.com/cattaneopablo/tuiter/handlers"
	"github.com/cattaneopablo/tuiter/db"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("connection is not established")
		return
	}

	handlers.Handlers()
}