package main

import "testing"

func TestTotalCoins(t *testing.T) {
	expected := float64(11)
	if observed := getUserCoins("aeiou"); observed != expected {
		t.Fatalf("totalCoins('aeiou') = %v, want = %v", observed, expected)
	}
}
