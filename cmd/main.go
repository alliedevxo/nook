package main

import (
	"log"
	nookdb "nook/internal/db"
)

func main() {
	db, err := nookdb.Open()
	if err != nil {
		log.Fatalf("open notebook db: %v", err)
	}

	if err := nookdb.InsertNote(db, 2, "My third note", "I love notes !!!"); err != nil {
		log.Fatalf("view notebook table: %v", err)
	}

	if err := nookdb.ViewNotes(db); err != nil {
		log.Fatalf("view notebook table: %v", err)
	}
}