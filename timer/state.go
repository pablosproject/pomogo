package timer

import "math"

type WorkState int

const (
	IDLE WorkState = iota
	WORK
	SHORTBREAK
	LONGBREAK
)

type State struct {
	State     WorkState
	WorkCount int
}

func NewState() *State {
	return &State{
		State:     IDLE,
		WorkCount: 0,
	}
}

func (s *State) Next() {
	var newState WorkState
	switch s.State {
	case IDLE, SHORTBREAK, LONGBREAK:
		s.WorkCount++
		newState = WORK
	case WORK:
		if int64(math.Mod(float64(s.WorkCount), 4)) == 0 {
			newState = LONGBREAK
		} else {
			newState = SHORTBREAK
		}
	}

	s.State = newState
}

func (s *State) Cancel() {
	s.State = IDLE
}
