package game

import (
	"fmt"
	"time"
)

type Game struct {
	session *Session
}

func NewGame() *Game {
	return &Game{
		session: NewSession("", ""),
	}
}

func (g *Game) Start() {
	fmt.Println("Starting a new typing speed game...")
	g.session.Start()
}

func (g *Game) TrackTypingSpeed(input string) {
	g.session.TrackInput(input)
}

func (g *Game) GetResults() (speed float64, duration time.Duration) {
	return g.session.CalculateSpeed()
}