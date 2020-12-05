package main

import (
	"bufio"
	"bytes"
	"testing"
)

type testCase struct {
	sample []byte
	route  slopeRoute
	want   int
}

var testSample = []byte(`..##.........##.........##.........##.........##.........##.......
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#`)

func TestThreePart1(t *testing.T) {
	var cases = []testCase{
		{
			sample: testSample,
			want:   7,
			route:  slopeRoute{3, 1},
		},
		{
			sample: []byte(`.........###......#...#.......#
.#.#...........#..#..#.........
#.......#.................#....`),
			want:  1,
			route: slopeRoute{3, 1},
		},
	}

	for _, c := range cases {
		assertEqual(t,
			doThree(bytes.NewReader(c.sample), []slopeRoute{{3, 1}}),
			c.want)
	}
}

func TestThreePart2(t *testing.T) {
	t.Run("multiScanner works", func(t *testing.T) {
		sample := []byte(`this
is
the
sample
!`)

		scanner := makeScanner(sample)
		scanFn := multiScanner(scanner, 2)

		ok := scanFn()
		assertEqual(t, ok, true)
		assertEqual(t, scanner.Text(), `is`)

		ok = scanFn()
		assertEqual(t, ok, true)
		assertNotNil(t, scanner)
		assertEqual(t, scanner.Text(), `sample`)

		ok = scanFn()
		assertEqual(t, ok, false)
	})

	t.Run("count trees works", func(t *testing.T) {

		cases := []testCase{
			{sample: testSample, want: 2, route: slopeRoute{1, 1}},
			{sample: testSample, want: 7, route: slopeRoute{3, 1}},
			{sample: testSample, want: 3, route: slopeRoute{5, 1}},
			{sample: testSample, want: 4, route: slopeRoute{7, 1}},
			{sample: testSample, want: 2, route: slopeRoute{1, 2}},
		}

		for _, c := range cases {
			got := countTreesOldSkool(makeScanner(c.sample), c.route)
			// got := countTreesScanning(makeScanner(c.sample), c.route)
			assertEqual(t, got, c.want)
		}
	})

	t.Run("whole thing works", func(t *testing.T) {
		routes := []slopeRoute{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}

		r := bytes.NewReader(testSample)
		got := doThree(r, routes)
		assertEqual(t, got, 336)
	})
}

func makeScanner(b []byte) *bufio.Scanner {
	return bufio.NewScanner(bytes.NewReader(b))
}
