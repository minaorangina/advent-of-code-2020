package main

import (
	"strings"
	"testing"
)

func TestSix(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	reader1 := strings.NewReader(input)
	reader2 := strings.NewReader(input)

	assertEqual(t, doSixPart1(reader1), 11)
	assertEqual(t, doSixPart2(reader2), 6)
}
