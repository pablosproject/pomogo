// TODO: check how to put this file in its own package to just export the needed references
// https://go.dev/doc/code
package main

import (
	"time"
)

type PomodoroConfig struct {
	workLenght       time.Duration
	shortPauseLenght time.Duration
	longPauseLenght  time.Duration
}

type PomodoroTimer struct {
	config    PomodoroConfig
	state     State
	ticker    *time.Ticker
	startTime time.Time
	StateC    chan WorkState
	TimeC     chan time.Duration
}

func NewTimer(config PomodoroConfig) *PomodoroTimer {
	timer := PomodoroTimer{
		config:    config,
		state:     State{},
		ticker:    time.NewTicker(200 * time.Millisecond),
		startTime: time.Time{},
		StateC:    make(chan WorkState, 1),
		TimeC:     make(chan time.Duration),
	}

	timer.init()
	return &timer
}

func (t *PomodoroTimer) init() {
	go func() {
		for range t.ticker.C {
			if t.state.State != IDLE {
				if t.remainingTime() <= 0 {
					t.state.Next()
					t.notifyState()
					t.resetTimer()
				} else {
					t.TimeC <- t.remainingTime()
				}
			}
		}
	}()
}

func (t *PomodoroTimer) Start() {
	t.resetTimer()
	t.state.Next()
	t.notifyState()
}

func (t *PomodoroTimer) Stop() {
	t.state.Cancel()
	t.notifyState()
}

func (t *PomodoroTimer) State() WorkState {
	return t.state.State
}

func (t *PomodoroTimer) remainingTime() time.Duration {
	workLenght := t.config.workLenght
	if t.state.State == SHORTBREAK {
		workLenght = t.config.shortPauseLenght
	}
	if t.state.State == LONGBREAK {
		workLenght = t.config.longPauseLenght
	}

	return (workLenght - time.Since(t.startTime)).Round(time.Millisecond)
}

func (t *PomodoroTimer) resetTimer() {
	t.startTime = time.Now()
}

func (t *PomodoroTimer) notifyState() {
	t.StateC <- t.State()
}
