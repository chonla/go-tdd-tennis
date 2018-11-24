package tennis

import "fmt"

type Scorer struct {
	a   int
	b   int
	end bool
}

func (s *Scorer) scoreRepresent(v int) string {
	rep := []string{"0", "15", "30", "40", "A"}
	return rep[v]
}

func (s *Scorer) Reset() string {
	s.a = 0
	s.b = 0
	s.end = false
	return s.currentScore()
}

func (s *Scorer) currentScore() string {
	return fmt.Sprintf("%s-%s", s.scoreRepresent(s.a), s.scoreRepresent(s.b))
}

func (s *Scorer) isPlayer1HasMatchPoint(player1Score, player2Score int) bool {
	if (player1Score == 3 && player2Score < 3) ||
		(player1Score == 4 && player2Score == 3) {
		return true
	}
	return false
}

func (s *Scorer) isPlayer1HasAdvantage(player1Score, player2Score int) bool {
	return player1Score == 4 && player2Score == 3
}

func (s *Scorer) AGetPoint() string {
	if s.end {
		return "Error"
	}
	if s.isPlayer1HasMatchPoint(s.a, s.b) {
		s.end = true
		return "A Wins"
	}
	if s.isPlayer1HasAdvantage(s.b, s.a) {
		s.b = 3
	} else {
		s.a++
	}
	return s.currentScore()
}

func (s *Scorer) BGetPoint() string {
	if s.end {
		return "Error"
	}
	if s.isPlayer1HasMatchPoint(s.b, s.a) {
		s.end = true
		return "B Wins"
	}
	if s.isPlayer1HasAdvantage(s.a, s.b) {
		s.a = 3
	} else {
		s.b++
	}
	return s.currentScore()
}
