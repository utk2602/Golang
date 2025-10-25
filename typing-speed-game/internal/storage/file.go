package storage

import (
	"encoding/json"
	"os"
	"sync"

	"typing-speed-game/pkg/models"
)

type FileStorage struct {
	mu sync.Mutex
	filePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{filePath: filePath}
}

func (fs *FileStorage) Save(records []models.Record) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(records)
}

func (fs *FileStorage) Load() ([]models.Record, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(fs.filePath)
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