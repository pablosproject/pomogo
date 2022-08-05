package main

import "testing"

func AssertState(t testing.TB, got State, want WorkState) {
	t.Helper()
	if got.State != want {
		t.Errorf("Got state %d, wanted %d", got.State, want)
	}
}

func TestState(t *testing.T) {
	t.Run("Test work loop (should alternate work and short break until long break)", func(t *testing.T) {
		state := NewState()

		AssertState(t, *state, IDLE)

		// 1st work
		state.Next()
		AssertState(t, *state, WORK)
		state.Next()
		AssertState(t, *state, SHORTBREAK)

		// 2nd work
		state.Next()
		AssertState(t, *state, WORK)
		state.Next()
		AssertState(t, *state, SHORTBREAK)

		// 3rd work
		state.Next()
		AssertState(t, *state, WORK)
		state.Next()
		AssertState(t, *state, SHORTBREAK)

		// 4th work
		state.Next()
		AssertState(t, *state, WORK)
		state.Next()
		AssertState(t, *state, LONGBREAK)

		// Back to work
		state.Next()
		AssertState(t, *state, WORK)
	})

	t.Run("Cancel during a short break return to IDLE", func(t *testing.T) {
		state := NewState()

		state.Next()
		AssertState(t, *state, WORK)
		state.Next()
		AssertState(t, *state, SHORTBREAK)

		state.Cancel()
		AssertState(t, *state, IDLE)
	})

	t.Run("Cancel during a long break return to IDLE", func(t *testing.T) {
		state := NewState()

		// 1st work
		state.Next()
		state.Next()

		// 2nd work
		state.Next()
		state.Next()

		// 3rd work
		state.Next()
		state.Next()

		// 4th work
		state.Next()
		state.Next()

		state.Cancel()
		AssertState(t, *state, IDLE)
	})

	t.Run("Cancel during a work return to IDLE", func(t *testing.T) {
		state := NewState()

		// 1st work
		state.Next()

		state.Cancel()
		AssertState(t, *state, IDLE)
	})

}
