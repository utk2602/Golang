package analytics

import "github.com/utk2602/typing-speed-game/pkg/models"

type TypingSpeedStats struct {
	AverageSpeed float64
	BestSpeed    float64
	WorstSpeed   float64
	TotalGames   int
	Records      []models.TypingSpeedRecord
}

func NewTypingSpeedStats() *TypingSpeedStats {
	return &TypingSpeedStats{
		AverageSpeed: 0,
		BestSpeed:    0,
		WorstSpeed:   0,
		TotalGames:   0,
		Records:      []models.TypingSpeedRecord{},
	}
}

func (s *TypingSpeedStats) AddRecord(record models.TypingSpeedRecord) {
	s.Records = append(s.Records, record)
	s.TotalGames++
	s.updateStats(record.Speed)
}

func (s *TypingSpeedStats) updateStats(speed float64) {
	if s.TotalGames == 1 {
		s.BestSpeed = speed
		s.WorstSpeed = speed
		s.AverageSpeed = speed
		return
	}

	if s.BestSpeed < speed {
		s.BestSpeed = speed
	}
	if s.WorstSpeed > speed || s.WorstSpeed == 0 {
		s.WorstSpeed = speed
	}
	s.AverageSpeed = (s.AverageSpeed*float64(s.TotalGames-1) + speed) / float64(s.TotalGames)
}

func (s *TypingSpeedStats) GetRecords() []models.TypingSpeedRecord {
	return s.Records
}
