package main

import "time"

type PomodoroTimer struct {
	workLenght       time.Duration
	shortPauseLenght time.Duration
	longPauseLenght  time.Duration
	ticker           *time.Ticker
	startTime        time.Time
}

func (t *PomodoroTimer) start() {
	t.ticker = time.NewTicker(200 * time.Millisecond)
	t.startTime = time.Now()
}

func (t *PomodoroTimer) stop() {
	t.ticker.Stop()
}

func (t *PomodoroTimer) remainingTime() time.Duration {
	return (t.workLenght - time.Since(t.startTime)).Round(time.Millisecond)
}
