package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func sevenPart1() int {
	reader := getReader("input/07.txt")
	return doSevenPart1(readerToString(reader), "shiny gold")
}

func sevenPart2() int {
	reader := getReader("input/07.txt")
	return doSevenPart2(readerToString(reader), "shiny gold")
}

func doSevenPart1(data []string, target string) int {
	return len(walk1(data, target, 0))
}

func doSevenPart2(data []string, target string) int {
	return walk2(data, target, 0)
}

func walk1(data []string, target string, depth int) map[string]struct{} {
	final := map[string]struct{}{}
	if len(data) == 0 {
		return final
	}

	bags := []string{}

	for _, line := range data {
		pattern := fmt.Sprintf("(?P<subject>.*) bags contain .* %s", target)
		r := regexp.MustCompile(pattern)

		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		for i, name := range r.SubexpNames() {
			if name != "" {
				final[matches[i]] = struct{}{}
				bags = append(bags, matches[i])
			}
		}
	}

	for _, b := range bags {
		res := walk1(data, b, depth+1)
		for key := range res {
			final[key] = struct{}{}
		}
	}

	return final
}

func walk2(data []string, target string, depth int) int {
	final := 0
	if len(data) == 0 {
		return final
	}

	bags := map[string]int{}

	for _, line := range data {
		pattern := fmt.Sprintf("%s bags contain (?P<stuff>[^no other].*)", target)
		r := regexp.MustCompile(pattern)

		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		var sections []string
		for i, name := range r.SubexpNames() {
			if name == "stuff" {
				if strings.Contains(matches[i], ",") {
					sections = strings.Split(matches[i], ",")
				} else {
					sections = []string{strings.Replace(matches[i], ".", "", -1)}
				}
			}
		}

		for _, s := range sections {
			numR := regexp.MustCompile("\\d+")
			numStr := numR.FindString(s)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}

			// extremely ugly...
			bag := strings.Replace(s, fmt.Sprintf("%d ", num), "", -1)
			bag = strings.Replace(bag, " bags", "", -1)
			bag = strings.Replace(bag, " bag", "", -1)
			bag = strings.Replace(bag, ".", "", -1)
			bag = strings.TrimRight(bag, " ")
			bag = strings.TrimLeft(bag, " ")

			bags[bag] = num
		}
	}

	for bag, num := range bags {
		res := walk2(data, bag, depth+1) * num
		final += num + res
	}

	return final
}
