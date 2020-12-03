package main

import "testing"

func TestOne(t *testing.T) {
	testInput := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	t.Run("part one works", func(t *testing.T) {
		want := 514579
		got := doOnePart1(testInput)

		assertEqual(t, got, want)
	})

	t.Run("part 2 works", func(t *testing.T) {
		want := 241861950
		got := doOnePart2(testInput)

		assertEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
