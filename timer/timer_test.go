package timer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testTimer() *PomodoroTimer {
	timer := NewTimer(PomodoroConfig{
		workLenght:       1 * time.Second,
		shortPauseLenght: 1 * time.Second,
		longPauseLenght:  1 * time.Second,
	})
	return timer
}

func TestTimerState(t *testing.T) {
	t.Run("Should receive work state on when starting", func(t *testing.T) {
		timer := testTimer()

		rcvState := make(chan WorkState)

		go func() {
			state := <-timer.StateC
			fmt.Print(state)
			rcvState <- state
		}()
		timer.Start()

		select {
		case state := <-rcvState:
			assert.Equal(t, WORK, state)
		case <-time.After(1500 * time.Millisecond):
			assert.Fail(t, "work state not received")
		}
	})

	t.Run("Should receive short break state after work", func(t *testing.T) {
		timer := testTimer()
		timer.Start()
		time.Sleep(1100 * time.Millisecond)

		select {
		case state := <-timer.StateC:
			assert.Equal(t, SHORTBREAK, state)
		case <-time.After(200 * time.Millisecond):
			assert.Fail(t, "short break state not received")
		}
	})

	// t.Run("Should receive long break after 8 sec", func(t *testing.T) {
	// 	timer := testTimer()
	// 	timer.Start()
	// 	time.Sleep(8100 * time.Millisecond)

	// 	for {
	// 		select {
	// 		case state := <-timer.StateC:
	// 			assert.Equal(t, LONGBREAK, state)
	// 		default:
	// 			assert.Fail(t, "Long Break state not received")
	// 		}
	// 	}
	// })

	// Test that after short time is passed I get a notification on state
	// Test that I recive time notifications
	// Test stop function during notification (I receive something)
	// Test that I do not receive time notification if stop is enabled
}
