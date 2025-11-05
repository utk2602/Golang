package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/utk2602/Golang/typing-speed-game/pkg/models"
)

type FileStorage struct {
	mu       sync.Mutex
	filePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{filePath: filePath}
}

func (fs *FileStorage) Save(records []models.TypingSpeedRecord) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(records)
}

func (fs *FileStorage) Load() ([]models.TypingSpeedRecord, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(fs.filePath)
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
