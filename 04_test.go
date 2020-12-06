package main

import (
	"strings"
	"testing"
)

var singleInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`

var testInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

var sample = `eyr:2029 byr:1931 hcl:z cid:128
ecl:amb hgt:150cm iyr:2015 pid:148714704

byr:2013 hgt:70cm pid:76982670 ecl:#4f9a1c
hcl:9e724b eyr:1981 iyr:2027

pid:261384974 iyr:2015
hgt:172cm eyr:2020
byr:2001 hcl:#59c2d9 ecl:amb cid:163`

func TestFourPart1(t *testing.T) {
	tt := []struct {
		name  string
		input string
		valid int
	}{
		{"single", singleInput, 1},
		{"multi", testInput, 2},
		{"sample", sample, 3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			assertEqual(t, doFour(reader), tc.valid)
		})
	}
}
