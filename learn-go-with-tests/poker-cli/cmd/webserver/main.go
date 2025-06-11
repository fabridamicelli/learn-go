package main

import (
	"log"
	"net/http"
	"poker-cli"
)

const dbFilename = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("problem serving on port 5000 %v", err)
	}
}
