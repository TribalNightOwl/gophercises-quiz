package main

import (
	"testing"
)

func TestParseLines(t *testing.T) {
	lines := [][]string{
		{"3+3", "6"},
		{"myquestion", "myanswer"},
		{"my long question", "my long answer"},
	}

	want := []problem{
		{"3+3", "6"},
		{"myquestion", "myanswer"},
		{"my long question", "my long answer"},
	}

	got := parseLines(lines)

	for i, v := range want {
		if v != got[i] {
			t.Errorf("want %q, got %q", v, got[i])
		}
	}
}
