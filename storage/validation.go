package storage

import (
	"log"
	"strings"
)

// this is for backup
// check if pre-defined data
func (s *Storage) ValidateData(id int, q string) bool {
	quoteString, err := s.GetData(id)
	if err != nil {
		log.Fatalln(err)
		return true
	}

	return strings.EqualFold(q, quoteString)
}
