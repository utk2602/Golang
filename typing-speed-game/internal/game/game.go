package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/utk2602/Golang/typing-speed-game/data"
	"github.com/utk2602/Golang/typing-speed-game/internal/analytics"
	"github.com/utk2602/Golang/typing-speed-game/internal/storage"
	"github.com/utk2602/Golang/typing-speed-game/internal/ui"
	"github.com/utk2602/Golang/typing-speed-game/pkg/models"
)

type Game struct {
	session   *Session
	storage   *storage.Storage
	analytics *analytics.Analytics
}

func NewGame() *Game {
	store := storage.NewStorage("typing_records.json")

	// Load existing records
	records, err := store.LoadRecords()
	analyticsInstance := analytics.NewAnalytics()
	if err == nil {
		for _, record := range records {
			analyticsInstance.AddRecord(record)
		}
	}

	return &Game{
		storage:   store,
		analytics: analyticsInstance,
	}
}

func (g *Game) StartNewSession() {
	playerName := ui.GetUserInput("Enter your name: ")

	// Select random text
	rand.Seed(time.Now().UnixNano())
	text := data.Texts[rand.Intn(len(data.Texts))]

	g.session = NewSession(playerName, text)

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Type the following text:")
	fmt.Println(text)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("\nPress ENTER when ready to start...")
	ui.GetUserInput("")

	g.session.Start()

	typedText := ui.GetUserInput("\nStart typing: ")
	g.session.EndSession(typedText)

	speed := g.session.CalculateTypingSpeed()
	accuracy := g.session.CalculateAccuracy()

	ui.DisplayResults(speed, accuracy)

	// Save record
	record := models.NewTypingSpeedRecord(playerName, speed)
	g.analytics.AddRecord(record)
	g.storage.SaveRecord(record)

	fmt.Println("\nRecord saved successfully!")
}

func (g *Game) DisplayAnalytics() {
	if len(g.analytics.Records) == 0 {
		fmt.Println("\nNo records found. Play a game first!")
		return
	}

	avgSpeed := g.analytics.AverageSpeed()
	bestSpeed := g.analytics.BestSpeed()
	worstSpeed := g.analytics.WorstSpeed()

	fmt.Println("\n" + strings.Repeat("=", 50))
	ui.DisplayAnalytics(avgSpeed, bestSpeed, worstSpeed)

	fmt.Printf("Total Games Played: %d\n", len(g.analytics.Records))

	fmt.Println("\nRecent Records:")
	start := len(g.analytics.Records) - 5
	if start < 0 {
		start = 0
	}

	for i := len(g.analytics.Records) - 1; i >= start; i-- {
		record := g.analytics.Records[i]
		fmt.Printf("  %s - %s: %.2f WPM\n",
			record.Timestamp.Format("2006-01-02 15:04:05"),
			record.PlayerName,
			record.Speed)
	}
	fmt.Println(strings.Repeat("=", 50))
}

func (g *Game) TrackTypingSpeed(input string) {
	if g.session != nil {
		g.session.TrackInput(input)
	}
}

func (g *Game) GetResults() (speed float64, duration time.Duration) {
	if g.session != nil {
		return g.session.CalculateSpeed()
	}
	return 0, 0
}
