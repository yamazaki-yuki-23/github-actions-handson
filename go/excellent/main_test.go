package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EventOrOdd(10)
	if result != "event" {
		t.Errorf("expected: event, actual: %s", result)
	}
}
