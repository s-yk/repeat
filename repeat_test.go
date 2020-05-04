package main

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	result := repeat("X", 1)
	if result != "X" {
		t.Errorf("expected: %s, but: %s", "X", result)
	}

	result = repeat("x", 3)
	if result != "xxx" {
		t.Errorf("expected: %s, but: %s", "xxx", result)
	}

	result = repeat("x", 0)
	if result != "" {
		t.Errorf("expected: %s, but: %s", "", result)
	}

	result = repeat("Ab3", 4)
	if result != "Ab3Ab3Ab3Ab3" {
		t.Errorf("expected: %s, but: %s", "Ab3Ab3Ab3Ab3", result)
	}

	result = repeat("あいうえお", 3)
	if result != "あいうえおあいうえおあいうえお" {
		t.Errorf("expected: %s, but: %s", "あいうえおあいうえおあいうえお", result)
	}

	result = repeat("あい1Aお", 2)
	if result != "あい1Aおあい1Aお" {
		t.Errorf("expected: %s, but: %s", "あい1Aおあい1Aお", result)
	}
}
