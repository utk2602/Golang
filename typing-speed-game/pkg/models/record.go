package models

import "time"

type TypingSpeedRecord struct {
	PlayerName string    `json:"player_name"`
	Speed      float64   `json:"speed"` // words per minute
	Timestamp  time.Time `json:"timestamp"`
}

func NewTypingSpeedRecord(playerName string, speed float64) TypingSpeedRecord {
	return TypingSpeedRecord{
		PlayerName: playerName,
		Speed:      speed,
		Timestamp:  time.Now(),
	}
}