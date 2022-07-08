package main

import (
	"github.com/cattaneopablo/tuiter/handlers"
	"github.com/cattaneopablo/tuiter/db"
)

func main() {
	if !db.checkConnection() {
		log.Fatal("connection is not established")
		return
	}

	handlers.Handlers()
}