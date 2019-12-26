package main

import "testing"

func TestGroupByName(t *testing.T) {
	names := []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
		"Emma", "Isabella", "Emily", "Madison",
		"Ava", "Olivia", "Sophia", "Abigail",
		"Elizabeth", "Chloe", "Samantha",
		"Addison", "Natalie", "Mia", "Alexis"}

	expected := []string{"Evan", "Neil", "Adam", "Matt", "Emma"}
	if observed := groupByLength(names); !testEq(observed[3], expected) {
		t.Fatalf("groupByLength(..) = %v, want = %v", observed[3], expected)
	}
}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
