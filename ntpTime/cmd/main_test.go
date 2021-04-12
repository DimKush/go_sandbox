package main

import "testing"

func TestCount(t *testing.T) {
	s := "Just simple string for count s"
	e := 4
	if c := Count(s, 's'); c != e {
		t.Fatalf("bad count for %s. got %d expected %d", s, c, e)
	}
}
