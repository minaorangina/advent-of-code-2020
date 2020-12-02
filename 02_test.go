package main

import (
	"strings"
	"testing"
)

func TestTwo(t *testing.T) {
	t.Run("parser works", func(t *testing.T) {
		input := []string{
			"1-3 a: abcde",
			"1-3 b: cdefg",
			"2-9 c: ccccccccc",
		}

		want := []passwordPolicy{
			{
				min:      1,
				max:      3,
				char:     "a",
				password: "abcde",
			},
			{
				min:      1,
				max:      3,
				char:     "b",
				password: "cdefg",
			},
			{
				min:      2,
				max:      9,
				char:     "c",
				password: "ccccccccc",
			},
		}

		for i, c := range input {
			assertEqual(t, parseLine(c), want[i])
		}
	})

	t.Run("charCountWithinRange works", func(t *testing.T) {
		input := `1-3 a: abcde
			1-3 b: cdefg"
			2-9 c: ccccccccc`
		r := strings.NewReader(input)
		assertEqual(t, sumValidPasswords(r, charCountWithinRange), 2)
	})

	t.Run("charPositionsValid works", func(t *testing.T) {
		input := `1-3 a: abcde
			1-3 b: cdefg"
			2-9 c: ccccccccc`
		r := strings.NewReader(input)
		assertEqual(t, sumValidPasswords(r, charPositionsValid), 1)
	})
}
