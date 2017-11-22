package main

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expected := 16

	if len(d) != expected {
		t.Errorf("Expected deck length of %v, but got %v", expected, len(d))
	}
}

