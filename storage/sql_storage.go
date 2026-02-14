package storage

import (
	"database/sql"
	"fmt"
	"strings"

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

// Get the whole data from database
func (s *Storage) GetWholeData() ([]core.Quote, error) {
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

// delete the data
func (s *Storage) DeleteData(id int) error {
	query := "DELETE FROM QUOTES WHERE id = ?"
	if _, err := s.DB.Exec(query, id); err != nil {
		return fmt.Errorf("failed to delete record: %v", err)
	}

	return nil
}

// delete the whole database
func (s *Storage) DeleteWholeData() error {
	query := "DELETE FROM QUOTES"
	if _, err := s.DB.Exec(query); err != nil {
		return fmt.Errorf("failed to delete whole record: %v", err)
	}

	return nil
}

// Batch insert
func (s *Storage) BatchInsert(quotes []core.Quote) error {
	var query strings.Builder

	query.WriteString("INSERT INTO QUOTES(id,quote,author) VALUES")
	var values []any

	for i, q := range quotes {
		if i > 0 {
			query.WriteString(",")
		}
		query.WriteString("(?,?,?)")
		values = append(values, q.Id, q.Quote, q.Author)
	}

	_, err := s.DB.Exec(query.String(), values...)
	return err
}

// get a selective data
func (s *Storage) GetData(id int) (string, error) {
	q := fmt.Sprintf("SELECT * FROM QUOTES WHERE id = %d", id)
	rows, err := s.DB.Query(q)
	if err != nil {
		return "", fmt.Errorf("no quote found with id: %v", err)
	}

	var quote core.Quote
	for rows.Next() {
		if err = rows.Scan(&quote.Id, &quote.Quote, &quote.Author); err != nil {
			return "", err
		}
	}

	return quote.Quote, nil
}
