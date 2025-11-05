package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/utk2602/typing-speed-game/pkg/models"
)

type Storage struct {
	mu       sync.Mutex
	filePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{filePath: filePath}
}

func (s *Storage) SaveRecord(record models.TypingSpeedRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.loadRecordsNoLock()
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
	encoder.SetIndent("", "  ")
	return encoder.Encode(records)
}

func (s *Storage) LoadRecords() ([]models.TypingSpeedRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.loadRecordsNoLock()
}

func (s *Storage) loadRecordsNoLock() ([]models.TypingSpeedRecord, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.TypingSpeedRecord{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var records []models.TypingSpeedRecord
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&records); err != nil {
		return []models.TypingSpeedRecord{}, nil
	}
	return records, nil
}
