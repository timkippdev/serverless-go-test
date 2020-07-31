package main_test

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log("Log from a test!")
	t.Fail("Simulate a test failure")
}
