// TODO: check how to put this file in its own package to just export the needed references
// https://go.dev/doc/code
package main

import (
	"math"
	"time"
)

type PomodoroState int

const (
	IDLE PomodoroState = iota
	WORK
	SHORTBREAK
	LONGBREAK
)

type PomodoroTimer struct {
	state            PomodoroState
	workCount        int
	workLenght       time.Duration
	shortPauseLenght time.Duration
	longPauseLenght  time.Duration
	ticker           *time.Ticker
	startTime        time.Time
	StateC           chan PomodoroState
	TimeC            chan time.Duration
}

func NewTimer(workLenght, shortBreak int) *PomodoroTimer {
	timer := PomodoroTimer{
		state:            IDLE,
		workCount:        0,
		workLenght:       time.Duration(workLenght) * time.Minute,
		shortPauseLenght: time.Duration(shortBreak) * time.Minute,
		longPauseLenght:  0,
		ticker:           time.NewTicker(200 * time.Millisecond),
		StateC:           make(chan PomodoroState, 1),
		TimeC:            make(chan time.Duration),
	}

	timer.workLenght = 3 * time.Second
	timer.shortPauseLenght = 3 * time.Second
	timer.longPauseLenght = 3 * time.Second

	timer.init()
	return &timer
}

func (t *PomodoroTimer) init() {
	go func() {
		for range t.ticker.C {
			if t.state != IDLE {
				if t.remainingTime() <= 0 {
					nextState := t.nextState()
					if nextState == WORK {
						t.workCount++
					}
					t.setState(nextState)
					t.resetTimer()
				} else {
					t.TimeC <- t.remainingTime()
				}
			}
		}
	}()
}

func (t *PomodoroTimer) Start() {
	t.workCount++
	t.setState(WORK)
	t.resetTimer()
}

func (t *PomodoroTimer) Stop() {
	t.workCount = 0
	t.setState(IDLE)
}

func (t *PomodoroTimer) State() PomodoroState {
	return t.state
}

func (t *PomodoroTimer) remainingTime() time.Duration {
	workLenght := t.workLenght
	if t.state == SHORTBREAK {
		workLenght = t.shortPauseLenght
	}
	if t.state == LONGBREAK {
		workLenght = t.longPauseLenght
	}

	return (workLenght - time.Since(t.startTime)).Round(time.Millisecond)
}

func (t *PomodoroTimer) setState(s PomodoroState) {
	t.state = s
	select {
	case t.StateC <- s:
	default:
	}
}

func (t *PomodoroTimer) resetTimer() {
	t.startTime = time.Now()
}

// TODO: to make things testable, extract the concept of state in a separate struct and test that
//		the sctruct must also include current pomodoro count for short/long break
func (t *PomodoroTimer) nextState() PomodoroState {
	var state PomodoroState
	switch t.state {
	case IDLE:
		state = WORK
	case WORK:
		if int64(math.Mod(float64(t.workCount), 4)) == 0 {
			state = LONGBREAK
		} else {
			state = SHORTBREAK
		}
	case SHORTBREAK, LONGBREAK:
		state = WORK
	default:
		state = IDLE
	}

	return state
}
