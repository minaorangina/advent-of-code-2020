package main

import (
	"bufio"
	"io"
	"strings"
)

func sixPart1() int {
	reader := getReader("input/06.txt")
	return doSixPart1(reader)
}

func sixPart2() int {
	reader := getReader("input/06.txt")
	return doSixPart2(reader)
}

func doSixPart1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitByBlankLine)

	count := 0

	for scanner.Scan() {
		// fmt.Println("ONE GROUP", scanner.Text())

		groupAnswers := map[string]struct{}{}
		perPerson := strings.Split(scanner.Text(), "\n")

		for _, answers := range perPerson {
			for _, c := range answers {
				groupAnswers[string(c)] = struct{}{}
			}
		}

		count += len(groupAnswers)
	}

	return count
}

func doSixPart2(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitByBlankLine)

	count := 0

	for scanner.Scan() {
		perPerson := strings.Split(scanner.Text(), "\n")
		groupSize := 0
		for _, c := range perPerson {
			if c != "" {
				groupSize++
			}
		}
		groupAnswers := map[string]int{}

		for _, answers := range perPerson {
			for _, c := range answers {
				groupAnswers[string(c)]++
			}
		}

		for _, c := range groupAnswers {
			if c == groupSize {
				count++
			}
		}
	}

	return count
}
