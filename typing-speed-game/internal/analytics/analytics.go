package analytics

import (
	"time"
)

type TypingSpeedRecord struct {
	PlayerName string
	Speed      float64 // words per minute
	Timestamp  time.Time
}

type Analytics struct {
	Records []TypingSpeedRecord
}

func NewAnalytics() *Analytics {
	return &Analytics{
		Records: []TypingSpeedRecord{},
	}
}

func (a *Analytics) AddRecord(record TypingSpeedRecord) {
	a.Records = append(a.Records, record)
}

func (a *Analytics) AverageSpeed() float64 {
	if len(a.Records) == 0 {
		return 0
	}

	var totalSpeed float64
	for _, record := range a.Records {
		totalSpeed += record.Speed
	}
	return totalSpeed / float64(len(a.Records))
}

func (a *Analytics) GetTrends() map[string]float64 {
	trends := make(map[string]float64)
	for _, record := range a.Records {
		date := record.Timestamp.Format("2006-01-02")
		trends[date] += record.Speed
	}
	return trends
}