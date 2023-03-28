package main

import (
	"log"
	"sre-cli/infrastructure/datastore"

	"sre-cli/registries"

	"sre-cli/infrastructure/cmd"
)

func main() {
	clients, err := datastore.NewClients()
	if err != nil {
		log.Panic(err)
	}

	r := registries.NewRegistry(clients)

	cmd.Execute(r.NewAppController())
}
