package main

import (
	"strings"
	"testing"
)

type q4Test struct {
	data  string
	valid bool
}

func TestFourPart1(t *testing.T) {
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

	tt := []struct {
		name  string
		data  string
		valid int
	}{
		{"single", singleInput, 1},
		{"multi", testInput, 2},
		{"sample", sample, 3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.data)
			assertEqual(t, doFour(reader, part1), tc.valid)
		})
	}
}

func TestFourPart2(t *testing.T) {
	badInput := `eyr:1972 cid:100
		hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926
		
		iyr:2019
		hcl:#602927 eyr:1967 hgt:170cm
		ecl:grn pid:012533040 byr:1946
		
		hcl:dab227 iyr:2012
		ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277
		
		hgt:59cm ecl:zzz
		eyr:2038 hcl:74454a iyr:2023
		pid:3556412378 byr:2007`

	goodInput := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

	oneGood := "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"

	_ = badInput
	_ = goodInput
	_ = oneGood

	t.Run("whole thing works", func(t *testing.T) {
		// t.Skip()
		tt := []struct {
			name     string
			data     string
			numValid int
		}{
			{"all bad", badInput, 0},
			{"all good", goodInput, 4},
			{"one good", oneGood, 1},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				reader := strings.NewReader(tc.data)
				assertEqual(t, doFour(reader, part2), tc.numValid)
			})
		}
	})

	t.Run("validateByr works", func(t *testing.T) {
		// t.Skip()
		tt := []q4Test{
			{"byr:2002", true},
			{"byr:2003", false},
			{"byr:1920", true},
			{"byr:1919", false},
			{"byr:191", false},
			{"byr:20022", false},
		}

		for _, tc := range tt {
			t.Run(tc.data, func(t *testing.T) {
				assertEqual(t, validateByr(tc.data), tc.valid)
			})
		}
	})

	t.Run("validateIyr works", func(t *testing.T) {
		tt := []q4Test{
			{"iyr:2010", true},
			{"iyr:2017", true},
			{"iyr:2020", true},
			{"iyr:2002", false},
			{"iyr:2021", false},
			{"iyr:1919", false},
			{"iyr:191", false},
			{"iyr:20100", false},
			{"iyr:20201", false},
		}

		for _, tc := range tt {
			t.Run(tc.data, func(t *testing.T) {
				assertEqual(t, validateIyr(tc.data), tc.valid)
			})
		}
	})

	t.Run("validateEyr works", func(t *testing.T) {
		tt := []q4Test{
			{"eyr:2020", true},
			{"eyr:2030", true},
			{"eyr:2021", true},
			{"eyr:2024", true},
			{"eyr:2002", false},
			{"eyr:1919", false},
			{"eyr:203", false},
			{"eyr:20100", false},
			{"eyr:20301", false},
		}

		for _, tc := range tt {
			t.Run(tc.data, func(t *testing.T) {
				assertEqual(t, validateEyr(tc.data), tc.valid)
			})
		}
	})

	t.Run("validateHgt works", func(t *testing.T) {
		tt := []q4Test{
			{"hgt:150cm", true},
			{"hgt:193cm", true},
			{"hgt:194cm", false},
			{"hgt:149cm", false},
			{"hgt:1500cm", false},
			{"hgt:1935cm", false},
			{"hgt:59in", true},
			{"hgt:76in", true},
			{"hgt:68in", true},
			{"hgt:58in", false},
			{"hgt:77in", false},
			{"hgt:588in", false},
			{"hgt:7in", false},
			{"hht:76in", false},
		}

		for _, tc := range tt {
			t.Run(tc.data, func(t *testing.T) {
				assertEqual(t, validateHgt(tc.data), tc.valid)
			})
		}
	})

	t.Run("validateHcl works", func(t *testing.T) {
		tt := []q4Test{
			{"hcl:#123abc", true},
			{"hcl:#123abz", false},
			{"hcl:123abc", false},
			{"#123abz", false},
			{"123abc", false},
		}

		for _, tc := range tt {
			t.Run(tc.data, func(t *testing.T) {
				assertEqual(t, validateHcl(tc.data), tc.valid)
			})
		}
	})

	t.Run("validateEcl works", func(t *testing.T) {
		tt := []q4Test{
			{"ecl:amb", true},
			{"ecl:blu", true},
			{"ecl:brn", true},
			{"ecl:gry", true},
			{"ecl:grn", true},
			{"ecl:hzl", true},
			{"ecl:oth", true},
			{"ecl:wat", false},
			{"ecl:ambb", false},
			{"ecl:blue", false},
			{"ecl:brnn", false},
			{"ecl:ggry", false},
			{"ecl:hzzll", false},
		}

		for _, tc := range tt {
			t.Run(tc.data, func(t *testing.T) {
				assertEqual(t, validateEcl(tc.data), tc.valid)
			})
		}
	})

	t.Run("validatePid works", func(t *testing.T) {
		tt := []q4Test{
			{"pid:000000001", true},
			{"pid:0123456789", false},
		}

		for _, tc := range tt {
			t.Run(tc.data, func(t *testing.T) {
				assertEqual(t, validatePid(tc.data), tc.valid)
			})
		}
	})
}
