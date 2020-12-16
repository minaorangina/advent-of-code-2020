package main

import (
	"fmt"
	"regexp"
)

func sevenPart1() int {
	reader := getReader("input/07.txt")
	return doSevenPart1(readerToString(reader), "shiny gold")
}

func doSevenPart1(data []string, target string) int {
	return len(walk(data, target, 0))
}

func walk(data []string, target string, depth int) map[string]struct{} {
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
		res := walk(data, b, depth+1)
		for key := range res {
			final[key] = struct{}{}
		}
	}

	return final
}
