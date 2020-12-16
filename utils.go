package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertNotNil(t *testing.T, got interface{}) {
	t.Helper()
	if got == nil {
		t.Error("unexpected nil")
	}
}

func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return file
}

func getReader(path string) io.ReadSeeker {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read file %q", err)
	}

	return strings.NewReader(string(b))
}

func readerToString(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
