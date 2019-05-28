package goodbye

import "testing"

func TestGoodbye(t *testing.T) {
	result := Goodbye(); if result != "Hello, world." {
		t.Errorf("Got: %s", result)
	}
}
