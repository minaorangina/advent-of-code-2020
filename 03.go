package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func threePart1() int {
	file, err := os.Open("input/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return doThreePart1(file)
}

func doThreePart1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	idx, treeCount, squareCount := 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if idx > len(line) {
			break
		}
		switch line[idx] {
		case '#':
			treeCount++
		case '.':
			squareCount++
		}
		idx = (idx + 3) % len(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return treeCount
}
