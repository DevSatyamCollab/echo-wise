package storage

import (
	"database/sql"
	"fmt"

	core "github.com/DevSatyamCollab/echo-wise/internal/core"
	_ "github.com/mattn/go-sqlite3"
)

// open the database
func (s *Storage) OpenDb() error {
	db, err := sql.Open("sqlite3", s.FileName)
	if err != nil {
		return err
	}

	// check if the connection is actuall alive
	if err := db.Ping(); err != nil {
		return err
	}

	s.DB = db
	return nil
}

// create the table
func (s *Storage) CreateTable() error {
	q := "CREATE TABLE IF NOT EXISTS QUOTES (id INTEGER PRIMARY KEY, quote TEXT NOT NULL, author TEXT)"
	if _, err := s.DB.Exec(q); err != nil {
		return err
	}

	return nil
}

// insert the data
func (s *Storage) AddData(q, a string) error {
	query := "INSERT INTO QUOTES (quote, author) VALUES (?,?)"
	if _, err := s.DB.Exec(query, q, a); err != nil {
		return err
	}

	return nil
}

// Get the data from database
func (s *Storage) GetData() ([]core.Quote, error) {
	var results []core.Quote

	q := "SELECT * FROM QUOTES;"
	rows, err := s.DB.Query(q)
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var q core.Quote
		if err := rows.Scan(&q.Id, &q.Quote, &q.Author); err != nil {
			return nil, err
		}

		results = append(results, q)
	}
	return results, nil
}

// update the data
func (s *Storage) UpdateData(id int, q, a string) error {
	query := "UPDATE QUOTES SET quote = ?, author = ? WHERE id == ?"

	res, err := s.DB.Exec(query, q, a, id)
	if err != nil {
		return fmt.Errorf("failed to update record: %v", err)
	}

	// check it row was actually found and updated
	if rowsAfftected, _ := res.RowsAffected(); rowsAfftected == 0 {
		return fmt.Errorf("no quote found with Id %d", id)
	}

	return nil
}
