package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type passwordPolicy struct {
	password, char string
	min, max       int
}

func charCountWithinRange(policy passwordPolicy) int {
	charCount := 0
	for _, char := range policy.password {
		if string(char) == policy.char {
			charCount++
		}
	}

	if charCount >= policy.min && charCount <= policy.max {
		return 1
	}

	return 0
}

func charPositionsValid(policy passwordPolicy) int {
	first, second :=
		string(policy.password[policy.min-1]), string(policy.password[policy.max-1])

	if (first == policy.char || second == policy.char) && first != second {
		return 1
	}
	return 0
}

func twoPart1() int {
	file, err := os.Open("input/02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return sumValidPasswords(file, charCountWithinRange)
}

func twoPart2() int {
	file, err := os.Open("input/02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return sumValidPasswords(file, charPositionsValid)
}

func sumValidPasswords(r io.Reader, validate func(policy passwordPolicy) int) int {
	scanner := bufio.NewScanner(r)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		policy := parseLine(line)
		result += validate(policy)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func parseLine(line string) passwordPolicy {
	r := regexp.MustCompile(
		`(?P<min>\d+)-(?P<max>\d+)\s(?P<char>[a-zA-Z])\:\s+(?P<password>[a-zA-Z]+)`)

	match := r.FindStringSubmatch(line)
	if len(match) == 0 {
		log.Fatal("bad entry")
	}

	groups := map[string]string{}
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			groups[name] = match[i]
		}
	}

	min, err := strconv.Atoi(groups["min"])
	if err != nil {
		log.Fatal("parsing error", err)
	}
	max, err := strconv.Atoi(groups["max"])
	if err != nil {
		log.Fatal("parsing error", err)
	}

	policy := passwordPolicy{
		password: groups["password"],
		char:     groups["char"],
		min:      min,
		max:      max,
	}

	return policy
}
