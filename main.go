package main

import (
	"log"

	"github.com/nahuelsv/twitter-backend/bd"
	"github.com/nahuelsv/twitter-backend/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Without connection to the Database")
	}
	handlers.Managers()
}
