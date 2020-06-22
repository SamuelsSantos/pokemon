package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	return m.Run()
}
