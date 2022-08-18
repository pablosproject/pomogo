package timer

import "math"

type WorkState int

const (
	IDLE WorkState = iota
	WORK
	SHORTBREAK
	LONGBREAK
)

type PomodoroState struct {
	state     WorkState
	workCount int
}

func NewState() *PomodoroState {
	return &PomodoroState{
		state:     IDLE,
		workCount: 0,
	}
}

func (s *PomodoroState) Next() {
	var newState WorkState
	switch s.state {
	case IDLE:
		s.workCount++
		newState = WORK
	case SHORTBREAK, LONGBREAK:
		newState = IDLE
	case WORK:
		if int64(math.Mod(float64(s.workCount), 4)) == 0 {
			newState = LONGBREAK
		} else {
			newState = SHORTBREAK
		}
	}

	s.state = newState
}

func (s *PomodoroState) Cancel() {
	s.state = IDLE
}
