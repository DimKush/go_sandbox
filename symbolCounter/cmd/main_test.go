package main

import (
	"testing"

	"github.com/DimKush/go_sandbox/tree/main/symbolCounter/internal/counter"
)

func TestUnpackStr(t *testing.T) {
	s := "a4bc2d5e"
	sExpected := "aaaabccddddde"

	str, err := counter.UnpackString(s)
	if err != nil {
		t.Fatalf("Error during counter.UnpackString(s):  %s", err.Error())
	}

	if str != sExpected || err != nil {
		t.Fatalf("Bad UnpackString for %s. Got %s expected %s", s, str, sExpected)
	}
}

func TestUnpackEscape(t *testing.T) {
	s := `qwe\\5`
	sExpected := `qwe\\\\\`

	str, err := counter.UnpackString(s)
	if err != nil {
		t.Fatalf("Error during counter.UnpackString(s):  %s", err.Error())
	}

	if str != sExpected || err != nil {
		t.Fatalf("Bad UnpackString for %s. Got %s expected %s", s, str, sExpected)
	}

	s = `qwe\4\5`
	sExpected = `qwe45`

	str, err = counter.UnpackString(s)
	if err != nil {
		t.Fatalf("Error during counter.UnpackString(s):  %s", err.Error())
	}

	if str != sExpected || err != nil {
		t.Fatalf("Bad UnpackString for %s. Got %s expected %s", s, str, sExpected)
	}
}

func TestCheckForLetters(t *testing.T) {
	str := "445"
	if err := counter.CheckForLetters(&str); err == nil {
		t.Fatalf("Bad UnpackString for %s. Expected an error", str)
	}

	str = "445e"
	if err := counter.CheckForLetters(&str); err != nil {
		t.Fatalf("Bad UnpackString for %s. Expected a err == nil", str)
	}
}

func TestPackLetters(t *testing.T) {
	s := "aaaabccddddde"
	strExpected := "a4bc2d5e"

	str, err := counter.PackLetters(s)
	if err != nil {
		t.Fatalf("Error during counter.PackLetters (s):  %s", err.Error())
	}

	if str != strExpected || err != nil {
		t.Fatalf("Bad PackLetters for %s. Got %s expected %s", s, str, strExpected)
	}
}
