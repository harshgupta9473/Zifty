package main

import (
	"log"

	"github.com/harshgupta9473/zifty/components/db"
	"github.com/harshgupta9473/zifty/components/workers"
)

func main() {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	err = store.Init()
	if err != nil {
		log.Fatal(err)
	}
	server := workers.NewServer(":3000",store)
	server.Run()
}
