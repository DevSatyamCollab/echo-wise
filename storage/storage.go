package storage

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	vaultDir string
	fn       = "quotes.db"
	once     sync.Once
	instance *Storage
)

type Storage struct {
	FileName string
	DB       *sql.DB
}

func InitializeStorage() *Storage {
	once.Do(func() {
		fname, err := unsureStorage()
		if err != nil {
			log.Fatalf("Initialize Failed:  %v", err)
		}

		instance = &Storage{
			FileName: fname,
		}
	})

	return instance
}

func unsureStorage() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}

	vaultDir = filepath.Join(homeDir, ".quotesApp")

	// create a directory
	err = os.MkdirAll(vaultDir, 0750)
	if err != nil {
		log.Fatalf("Error can't create directory: %v", err)
	}

	fname := filepath.Join(vaultDir, fn)

	// sqlite3 handles the file creation
	return fname, nil
}
