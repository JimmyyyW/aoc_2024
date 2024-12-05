package day5

import "testing"

func TestParseInput(t *testing.T) {
	pageMap, updates := parseInput("input.txt")
	if len(pageMap) != 49 {
		t.Fatalf("length should be xxx but was %d", len(pageMap))
	}
	if len(updates) != 199 {
		t.Fatalf("length should be xxx but was %d", len(updates))
	}
}

func TestPartOneExample(t *testing.T) {
	pageMap, updates := parseInput("example.txt")
	result, bad := solve(pageMap, updates)
	if result != 143 {
		t.Fatalf("expected 143 but was %d", result)
	}
	if bad != 123 {
		t.Fatalf("for bad - expected 123 but was %d", bad)
	}
}

func TestPartOne(t *testing.T) {
	pageMap, updates := parseInput("input.txt")
	result, bad := solve(pageMap, updates)
	if result != 5713 {
		t.Fatalf("expected xx but was %d", result)
	}
	if bad != 5713 {
		t.Fatalf("bad - expected xx but was %d", bad)
	}
}
