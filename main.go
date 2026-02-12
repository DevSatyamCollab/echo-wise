package main

import (
	"log"
	"os"

	"github.com/DevSatyamCollab/echo-wise/storage"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store := storage.InitializeStorage()
	// open db
	if err := store.OpenDb(); err != nil {
		log.Fatalf("Error failed to open database: %v", err)
	}

	// create a table
	if err := store.CreateTable(); err != nil {
		log.Fatalf("Error can't create table in Database: %v", err)
	}

	// close db
	defer func() {
		if err := store.DB.Close(); err != nil {
			log.Fatalf("Error failed to close database: %v", err)
		}
	}()

	p := tea.NewProgram(InitialModel(store), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
		os.Exit(1)
	}
}
