package main

import (
	"strings"
	"testing"
)

func TestFivePart1(t *testing.T) {
	t.Run("getSeatID works", func(t *testing.T) {
		assertEqual(t, getSeatID(44, 5), 357)
	})

	t.Run("getMaxInt works", func(t *testing.T) {
		got := getMaxInt([]int{4, 7, 2, 5, 6, 2, 7, 3, 4, 75, 434, 5, 2})
		assertEqual(t, got, 434)
	})

	t.Run("getRowWorks", func(t *testing.T) {
		tt := []struct {
			code string
			want int
		}{
			{"FBFBBFF", 44},
			{"BFFFBBF", 70},
			{"FFFBBBF", 14},
			{"BBFFBBF", 102},
		}

		for _, tc := range tt {
			t.Run(tc.code, func(t *testing.T) {
				got := getRow([]byte(tc.code), 0, maxRow)
				assertEqual(t, got, tc.want)
			})
		}
	})

	t.Run("getColumn works", func(t *testing.T) {
		tt := []struct {
			code string
			want int
		}{
			{"RLR", 5},
			{"RRR", 7},
			{"RLL", 4},
		}

		for _, tc := range tt {
			t.Run(tc.code, func(t *testing.T) {
				got := getColumn([]byte(tc.code), 0, maxColumn)
				assertEqual(t, got, tc.want)
			})
		}
	})

	t.Run("whole thing works", func(t *testing.T) {
		input := `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`
		reader := strings.NewReader(input)
		assertEqual(t, doFivePart1(reader), 820)
	})
}
