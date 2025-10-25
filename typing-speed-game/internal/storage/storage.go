package storage

import (
	"encoding/json"
	"os"
	"sync"

	"typing-speed-game/pkg/models"
)

type Storage struct {
	mu sync.Mutex
	filePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{filePath: filePath}
}

func (s *Storage) SaveRecord(record models.Record) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.LoadRecords()
	if err != nil {
		return err
	}

	records = append(records, record)

	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(records)
}

func (s *Storage) LoadRecords() ([]models.Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Record{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var records []models.Record
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&records)
	return records, err
}