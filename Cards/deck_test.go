package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 13*4 {
		// t.Error(d, "Expected deck sze of 52, but got ", len(d))
		t.Errorf("Expected deck sze of 52, but got %v", len(d))
	}
}
