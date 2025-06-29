package main

import (
	"log"
	"net/http"
	"os"
)

const dbFilename = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 06666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("problem serving on port 5000 %v", err)
	}
}

//TODO https://quii.gitbook.io/learn-go-with-tests/build-an-application/io#another-problem
