package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func fourPart1() int {
	b, err := ioutil.ReadFile("input/04.txt")
	if err != nil {
		log.Fatalf("could not read file %q", err)
	}
	reader := strings.NewReader(string(b))

	return doFour(reader, part1)
}

func fourPart2() int {
	b, err := ioutil.ReadFile("input/04.txt")
	if err != nil {
		log.Fatalf("could not read file %q", err)
	}
	reader := strings.NewReader(string(b))

	return doFour(reader, part2)
}

func doFour(r io.Reader, validator func(s string) int) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitByBlankLine)

	count := 0
	for scanner.Scan() {
		replacer := strings.NewReplacer("\n", " ")
		trimmed := replacer.Replace(scanner.Text())

		numValid := validator(trimmed)
		if numValid == 7 {
			count++
		}
	}

	return count
}

func part1(data string) int {
	requiredFields := map[string]struct{}{
		"byr": struct{}{},
		"iyr": struct{}{},
		"ecl": struct{}{},
		"eyr": struct{}{},
		"hgt": struct{}{},
		"hcl": struct{}{},
		"pid": struct{}{},
	}

	r := regexp.MustCompile(":\\S+")
	fields := r.Split(strings.TrimSpace(data), -1)

	validFields := 0
	for _, f := range fields {
		fld := strings.TrimSpace(f)
		if _, ok := requiredFields[fld]; ok {
			validFields++
		}
	}

	return validFields
}

func part2(data string) int {
	requiredFields := map[string]func(string) bool{
		"byr": validateByr,
		"iyr": validateIyr,
		"ecl": validateEcl,
		"eyr": validateEyr,
		"hgt": validateHgt,
		"hcl": validateHcl,
		"pid": validatePid,
	}

	r := regexp.MustCompile(":\\S+")
	fields := r.Split(strings.TrimSpace(data), -1)

	validFields := 0
	for _, f := range fields {
		f = strings.TrimSpace(f)
		fn, ok := requiredFields[f]
		if !ok {
			continue
		}
		if valid := fn(data); valid {
			validFields++
		}
	}

	return validFields
}

func doNothing(s string) bool {
	return true
}

func validateByr(s string) bool {
	r := regexp.MustCompile("byr:[0-9]{4}(\\s|$)")
	pair := r.FindString(s)
	if pair == "" {
		return false
	}
	data := strings.Trim(pair, "byr: ")
	if len(data) != 4 {
		return false
	}
	n, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if 1920 > n || 2002 < n {
		return false
	}
	return true
}

func validateEcl(s string) bool {
	r := regexp.MustCompile("ecl:[a-zA-Z]{3}(\\s|$)")
	match := r.FindString(s)
	if match == "" {
		return false
	}

	data := strings.TrimPrefix(match, "ecl:")

	ws := regexp.MustCompile("(\\s|$)")
	data = ws.ReplaceAllString(data, "")
	if len(data) != 3 {
		return false
	}

	allowedColours := map[string]struct{}{
		"amb": struct{}{},
		"blu": struct{}{},
		"brn": struct{}{},
		"gry": struct{}{},
		"grn": struct{}{},
		"hzl": struct{}{},
		"oth": struct{}{},
	}

	if _, ok := allowedColours[data]; !ok {
		return false
	}

	return true
}

func validateEyr(s string) bool {
	r := regexp.MustCompile("eyr:[0-9]{4}(\\s|$)")
	pair := r.FindString(s)
	if pair == "" {
		return false
	}
	data := strings.Trim(pair, "eyr: ")
	if len(data) != 4 {
		return false
	}
	n, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if 2020 > n || 2030 < n {
		return false
	}
	return true
}

func validateHcl(s string) bool {
	r := regexp.MustCompile("hcl:#[A-Fa-f0-9]{6}")
	match := r.FindString(s)
	if match == "" {
		return false
	}
	return true
}

func validateHgt(s string) bool {
	r := regexp.MustCompile("hgt:(?P<value>[0-9]{2,3})(?P<units>(cm)|(in))")
	match := r.FindString(s)
	if match == "" {
		return false
	}

	parts := r.FindStringSubmatch(s)
	idx := r.SubexpIndex("value")
	if parts[idx] == "" {
		return false
	}

	rawValue := parts[idx]
	if len(rawValue) > 3 {
		return false
	}

	idx = r.SubexpIndex("units")
	if parts[idx] == "" {
		return false
	}
	units := parts[idx]

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return false
	}

	switch units {
	case "cm":
		if value < 150 || value > 193 {
			return false
		}
	case "in":
		if value < 59 || value > 76 {
			return false
		}
	default:
		return false
	}

	return true
}

func validateIyr(s string) bool {
	r := regexp.MustCompile("iyr:[0-9]{4}(\\s|$)")
	pair := r.FindString(s)
	if pair == "" {
		return false
	}
	data := strings.Trim(pair, "iyr: ")
	if len(data) != 4 {
		return false
	}
	n, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if 2010 > n || 2020 < n {
		return false
	}
	return true
}

func validatePid(s string) bool {
	r := regexp.MustCompile("pid:[0-9]{9}(\\s|$)")
	match := r.FindString(s)
	if match == "" {
		return false
	}

	data := strings.TrimPrefix(match, "pid:")
	ws := regexp.MustCompile("\\s|$")
	data = ws.ReplaceAllString(data, "")
	if len(data) != 9 {
		return false
	}

	if _, err := strconv.Atoi(data); err != nil {
		return false
	}

	return true
}

func splitByBlankLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
