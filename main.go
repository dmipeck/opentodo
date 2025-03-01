package main

import (
	"log"

	"gitlab.build13.com/opentodo/opentodo/internal/server"
)

func main() {
	log.Print("running server on :80...")
	err := server.Run(":80")
	if err != nil {
		panic(err)
	}
	log.Print("server stopped")
}
