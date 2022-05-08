// MIT License
// Copyright (c) 2022 Brian Reece

package hashmap

import "testing"

const (
	EXPECTED string = "expected %s (%d), got (%d)\n"
	ERROR    string = "experienced error: \"%s\"\n"
)

func setup() (*HashMap[string, int], map[string]int) {
	hashmap := New[string, int]()
	pairs := make(map[string]int, 5)
	pairs["apple"] = 1
	pairs["banana"] = 2
	pairs["cucumber"] = 3
	pairs["date"] = 4
	pairs["eclair"] = 5
	return hashmap, pairs
}

func TestNew(t *testing.T) {
	hashmap := New[string, int]()
	if len(*hashmap.Vector) != 0 {
		t.Errorf(EXPECTED, "len", 0, len(*hashmap.Vector))
	}
}