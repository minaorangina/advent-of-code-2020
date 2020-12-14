package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
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
