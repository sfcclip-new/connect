package main

import "testing"

func TestNewHashID(t *testing.T) {
	hashID := newHashID()
	if len(hashID) == 0 {
		t.Fatalf("Generated HashID is too short")
	}
}
