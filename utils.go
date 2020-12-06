package main

import (
	"log"
	"os"
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
