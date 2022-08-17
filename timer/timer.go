package timer

import (
	"time"
)

type PomodoroConfig struct {
	WorkLenght       time.Duration
	ShortPauseLenght time.Duration
	LongPauseLenght  time.Duration
}

type PomodoroTimer struct {
	config    PomodoroConfig
	state     State
	ticker    *time.Ticker
	startTime time.Time
}

func NewTimer(config PomodoroConfig) *PomodoroTimer {
	timer := PomodoroTimer{
		config:    config,
		state:     State{},
		ticker:    time.NewTicker(200 * time.Millisecond),
		startTime: time.Time{},
	}

	timer.init()
	return &timer
}

func (t *PomodoroTimer) init() {
	go func() {
		for range t.ticker.C {
			if t.state.State != IDLE {
				if t.RemainingTime() <= 0 {
					t.state.Next()
					t.resetTimer()
				}
			}
		}
	}()
}

func (t *PomodoroTimer) Start() {
	t.resetTimer()
	t.state.Next()
}

func (t *PomodoroTimer) Stop() {
	t.state.Cancel()
}

func (t *PomodoroTimer) State() WorkState {
	return t.state.State
}

func (t *PomodoroTimer) RemainingTime() time.Duration {
	workLenght := t.config.WorkLenght
	if t.state.State == SHORTBREAK {
		workLenght = t.config.ShortPauseLenght
	}
	if t.state.State == LONGBREAK {
		workLenght = t.config.LongPauseLenght
	}

	return (workLenght - time.Since(t.startTime)).Round(time.Millisecond)
}

func (t *PomodoroTimer) resetTimer() {
	t.startTime = time.Now()
}
