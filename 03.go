package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type slopeRoute struct {
	right, down int
}

func threePart1() int {
	file, err := os.Open("input/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return doThree(file, []slopeRoute{{3, 1}})
}

func threePart2() int {
	file, err := os.Open("input/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return doThree(file, []slopeRoute{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	})
}

func doThree(r io.ReadSeeker, routes []slopeRoute) int {
	// counts := []int{}
	result := 1
	for _, rt := range routes {
		count := countTreesOldSkool(bufio.NewScanner(r), rt)
		// count := countTreesScanning(bufio.NewScanner(r), rt)
		result = result * count
		r.Seek(0, 0)
	}

	return result
}

func countTreesOldSkool(scanner *bufio.Scanner, route slopeRoute) int {
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	idx, treeCount := route.right, 0
	for i := route.down; i < len(lines); i += route.down {
		switch lines[i][idx] {
		case '#':
			treeCount++
		}
		idx = (idx + route.right) % len(lines[i])
	}

	return treeCount
}

func countTreesScanning(scanner *bufio.Scanner, route slopeRoute) int {
	myScanner := multiScanner(scanner, route.down)

	idx, treeCount := route.right, 0
	// myScanner()
	for myScanner() {
		line := scanner.Text()
		if idx > len(line) {
			break
		}
		switch line[idx] {
		case '#':
			treeCount++
		}
		idx = (idx + route.right) % len(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return treeCount
}

func multiScanner(scanner *bufio.Scanner, down int) func() bool {
	fn := func() bool {
		var ok bool
		for i := 0; i < down; i++ {
			ok = scanner.Scan()
		}
		return ok
	}
	return fn
}
