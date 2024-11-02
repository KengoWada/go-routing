package main

import (
	"log"

	"github.com/KengoWada/gorouting/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8000")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
