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

func (s *Session) CalculateSpeed() (float64, time.Duration) {
	panic("unimplemented")
}

func (s *Session) TrackInput(input string) {
	panic("unimplemented")
}

func (s *Session) Start() {
	panic("unimplemented")
}

func NewSession(playerName string, inputText string) *Session {
	return &Session{
		PlayerName: playerName,
		StartTime:  time.Now(),
		InputText:  inputText,
	}
}

func (s *Session) EndSession(typedText string) {
	s.TypedText = typedText
	s.EndTime = time.Now()
}

func (s *Session) CalculateTypingSpeed() float64 {
	duration := s.EndTime.Sub(s.StartTime).Seconds()
	wordsTyped := float64(len(s.TypedText)) / 5.0 // Average word length
	typingSpeed := wordsTyped / duration * 60.0   // Words per minute
	return typingSpeed
}

func (s *Session) IsInputCorrect() bool {
	return s.InputText == s.TypedText
}
