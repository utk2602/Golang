package analytics

import (
	"github.com/utk2602/Golang/typing-speed-game/pkg/models"
)

type Analytics struct {
	Records []models.TypingSpeedRecord
}

func NewAnalytics() *Analytics {
	return &Analytics{
		Records: []models.TypingSpeedRecord{},
	}
}

func (a *Analytics) AddRecord(record models.TypingSpeedRecord) {
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

func (a *Analytics) BestSpeed() float64 {
	if len(a.Records) == 0 {
		return 0
	}

	best := a.Records[0].Speed
	for _, record := range a.Records {
		if record.Speed > best {
			best = record.Speed
		}
	}
	return best
}

func (a *Analytics) WorstSpeed() float64 {
	if len(a.Records) == 0 {
		return 0
	}

	worst := a.Records[0].Speed
	for _, record := range a.Records {
		if record.Speed < worst {
			worst = record.Speed
		}
	}
	return worst
}

func (a *Analytics) GetTrends() map[string]float64 {
	trends := make(map[string]float64)
	counts := make(map[string]int)

	for _, record := range a.Records {
		date := record.Timestamp.Format("2006-01-02")
		trends[date] += record.Speed
		counts[date]++
	}

	// Calculate average for each day
	for date := range trends {
		trends[date] = trends[date] / float64(counts[date])
	}

	return trends
}
