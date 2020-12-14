package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const (
	maxRow    = 127
	maxColumn = 7
)

func fivePart1() int {
	b, err := ioutil.ReadFile("input/05.txt")
	if err != nil {
		log.Fatalf("could not read file %q", err)
	}
	reader := strings.NewReader(string(b))

	return doFivePart1(reader)
}

func fivePart2() int {
	b, err := ioutil.ReadFile("input/05.txt")
	if err != nil {
		log.Fatalf("could not read file %q", err)
	}
	reader := strings.NewReader(string(b))

	return doFivePart2(reader)
}

func doFivePart1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var highest int
	for scanner.Scan() {
		code := scanner.Text()
		if len(code) != 10 {
			log.Fatalf("invalid code %s", code)
		}

		row := getRow([]byte(code)[:7], 0, maxRow)
		col := getColumn([]byte(code)[7:], 0, maxColumn)

		id := getSeatID(row, col)

		if id > highest {
			highest = id
		}
	}

	return highest
}

func doFivePart2(r io.Reader) int {
	tickets := map[int]struct{}{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		code := scanner.Text()
		if len(code) != 10 {
			log.Fatalf("invalid code %s", code)
		}

		row := getRow([]byte(code)[:7], 0, maxRow)
		col := getColumn([]byte(code)[7:], 0, maxColumn)

		id := getSeatID(row, col)
		tickets[id] = struct{}{}
	}

	expected := map[int]struct{}{}
	firstRow := map[int]struct{}{}
	lastRow := map[int]struct{}{}

	for c := 0; c <= maxColumn; c++ {
		for r := 0; r <= maxRow; r++ {
			id := getSeatID(r, c)
			expected[id] = struct{}{}

			if r == 0 {
				firstRow[id] = struct{}{}
			}
			if r == maxRow {
				lastRow[id] = struct{}{}
			}
		}
	}

	candidates := []int{}
	for id := range expected {
		if _, ok1 := tickets[id]; !ok1 {
			if _, ok2 := firstRow[id]; !ok2 {
				if _, ok3 := lastRow[id]; !ok3 {
					candidates = append(candidates, id)
				}
			}
		}
	}

	sort.Ints(candidates)

	mySeat := []int{}
	for i, v := range candidates {
		if i == 0 || i == len(candidates)-1 {
			continue
		}
		if v-candidates[i-1] > 1 && candidates[i+1]-v > 1 {
			mySeat = append(mySeat, candidates[i])
		}
	}

	if len(mySeat) != 1 {
		log.Fatalf("got more than one seat: %+v", mySeat)
	}
	return mySeat[0]
}

func getColumn(code []byte, low, high int) int {
	if len(code) == 0 {
		return low
	}
	newLow, newHigh := low, high
	for _, c := range code {
		char := string(c)

		if char == "L" {
			newHigh = low + (high-low)/2
		}
		if char == "R" {
			newLow = low + ((high - low) / 2) + 1
		}

		return getColumn(code[1:], newLow, newHigh)
	}

	return 0
}

func getRow(code []byte, low, high int) int {
	if len(code) == 0 {
		return low
	}
	newLow, newHigh := low, high

	for _, c := range code {
		char := string(c)
		if char == "F" {
			newHigh = low + (high-low)/2
		}
		if char == "B" {
			newLow = low + ((high - low) / 2) + 1
		}

		return getRow(code[1:], newLow, newHigh)
	}

	return 0
}

func getSeatID(row, column int) int {
	return (row * 8) + column
}

func getMaxInt(s []int) int {
	var largest int
	for _, v := range s {
		if v > largest {
			largest = v
		}
	}
	return largest
}
