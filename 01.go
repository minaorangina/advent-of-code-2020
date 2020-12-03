package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var targetValue = 2020

func onePart1() int {
	return doOnePart1(getData())
}

func onePart2() int {
	return doOnePart2(getData())
}

func doOnePart1(values []int) int {
	for _, val1 := range values {
		for _, val2 := range values {
			if val1+val2 == targetValue {
				return val1 * val2
			}
		}
	}

	return 0
}
func doOnePart2(values []int) int {
	for _, val1 := range values {
		for _, val2 := range values {
			for _, val3 := range values {
				if val1+val2+val3 == targetValue {
					return val1 * val2 * val3
				}
			}
		}
	}

	return 0
}

func getData() []int {
	file, err := os.Open("input/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	values := []int{}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}
