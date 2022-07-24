// MIT License
// Copyright (c) 2022 Brian Reece

package hash

import "testing"

const (
	EXPECTED string = "expected %s (%d), got (%d)\n"
	ERROR    string = "experienced error: \"%v\"\n"
)

func setup() (*Map[string, int], map[string]int) {
	hashmap := New[string, int](0.75, 5)
	pairs := make(map[string]int, 5)
	pairs["apple"] = 1
	pairs["banana"] = 2
	pairs["cucumber"] = 3
	pairs["date"] = 4
	pairs["eclair"] = 5
	return hashmap, pairs
}

func TestNew(t *testing.T) {
	hashmap := New[string, int](0.75, 5)
	if len(*hashmap.Vector) != 5 {
		t.Errorf(EXPECTED, "len", 5, len(*hashmap.Vector))
	}
}

func TestGet(t *testing.T) {
	hashmap, pairs := setup()
	got, err := hashmap.Get("cucumber")
	expected := pairs["cucumber"]
	if err != nil {
		t.Errorf(ERROR, err)
	}
	if got == nil || *got != expected {
		t.Errorf(EXPECTED, "value", expected, got)
	}
}

func TestSet(t *testing.T) {
	hashmap := New[string, int](0.75, 5)
	err := hashmap.Set("banana", 2)
	if err != nil {
		t.Errorf(ERROR, err)
	}
	got, err := hashmap.Get("banana")
	if err != nil {
		t.Errorf(ERROR, err)
	}
	if *got != 2 {
		t.Errorf(EXPECTED, "value", 2, *got)
	}
}
