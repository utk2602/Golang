package game

import (
	"time"
)

type Session struct {
	PlayerName string
	StartTime  time.Time
	EndTime    time.Time
	InputText  string
	TypedText  string
}

func NewSession(playerName string, inputText string) *Session {
	return &Session{
		PlayerName: playerName,
		InputText:  inputText,
	}
}

func (s *Session) Start() {
	s.StartTime = time.Now()
}

func (s *Session) TrackInput(input string) {
	s.TypedText = input
}

func (s *Session) EndSession(typedText string) {
	s.TypedText = typedText
	s.EndTime = time.Now()
}

func (s *Session) CalculateTypingSpeed() float64 {
	duration := s.EndTime.Sub(s.StartTime).Minutes()
	if duration == 0 {
		duration = 0.01
	}
	wordsTyped := float64(len(s.TypedText)) / 5.0 // Average word length
	typingSpeed := wordsTyped / duration          // Words per minute
	return typingSpeed
}

func (s *Session) CalculateSpeed() (float64, time.Duration) {
	duration := s.EndTime.Sub(s.StartTime)
	wordsTyped := float64(len(s.TypedText)) / 5.0 // Average word length
	minutes := duration.Minutes()
	if minutes == 0 {
		minutes = 0.01 // Prevent division by zero
	}
	typingSpeed := wordsTyped / minutes // Words per minute
	return typingSpeed, duration
}

func (s *Session) CalculateAccuracy() float64 {
	if len(s.InputText) == 0 {
		return 0
	}

	correctChars := 0
	minLen := len(s.InputText)
	if len(s.TypedText) < minLen {
		minLen = len(s.TypedText)
	}

	for i := 0; i < minLen; i++ {
		if s.InputText[i] == s.TypedText[i] {
			correctChars++
		}
	}

	return (float64(correctChars) / float64(len(s.InputText))) * 100
}

func (s *Session) IsInputCorrect() bool {
	return s.InputText == s.TypedText
}
