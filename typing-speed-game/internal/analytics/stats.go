package analytics

type TypingSpeedStats struct {
	AverageSpeed float64
	BestSpeed    float64
	TotalGames   int
	Records      []TypingSpeedRecord
}

type TypingSpeedRecord struct {
	PlayerName string
	Speed      float64 // words per minute
	Timestamp  int64   // Unix timestamp
}

func NewTypingSpeedStats() *TypingSpeedStats {
	return &TypingSpeedStats{
		AverageSpeed: 0,
		BestSpeed:    0,
		TotalGames:   0,
		Records:      []TypingSpeedRecord{},
	}
}

func (s *TypingSpeedStats) AddRecord(record TypingSpeedRecord) {
	s.Records = append(s.Records, record)
	s.TotalGames++
	s.updateStats(record.Speed)
}

func (s *TypingSpeedStats) updateStats(speed float64) {
	if s.BestSpeed < speed {
		s.BestSpeed = speed
	}
	s.AverageSpeed = (s.AverageSpeed*float64(s.TotalGames-1) + speed) / float64(s.TotalGames)
}

func (s *TypingSpeedStats) GetRecords() []TypingSpeedRecord {
	return s.Records
}