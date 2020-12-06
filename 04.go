package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func fourPart1() int {
	b, err := ioutil.ReadFile("input/04.txt")
	if err != nil {
		log.Fatalf("could not read file %q", err)
	}
	reader := strings.NewReader(string(b))

	return doFour(reader)
}

func doFour(r io.Reader) int {
	requiredFields := map[string]struct{}{
		"byr": struct{}{},
		"iyr": struct{}{},
		"ecl": struct{}{},
		"eyr": struct{}{},
		"hgt": struct{}{},
		"hcl": struct{}{},
		"pid": struct{}{},
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(customSplitFn)

	count := 0
	for scanner.Scan() {
		replacer := strings.NewReplacer("\n", " ")
		trimmed := replacer.Replace(scanner.Text())

		r := regexp.MustCompile(":\\S+")
		fields := r.Split(strings.TrimSpace(trimmed), -1)

		validFields := 0
		for _, f := range fields {
			fld := strings.TrimSpace(f)
			if _, ok := requiredFields[fld]; ok {
				validFields++
			}
		}
		if validFields == 7 {
			count++
		}
	}

	return count
}

func customSplitFn(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
