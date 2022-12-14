package main

import (
	"testing"
)

func TestReadFileLinesReturnsFirstLine(t *testing.T) {
	have, err := readFileLines("data.txt")
	want := "Lorem Ipsum is simply dummy text of the printing and typesetting"

	if err != nil {
		t.Errorf("have %#q", err)
	}

	if (have[0] != want) {
		t.Errorf("have \"%v\" want \"%v\"", have[0], want)
	}
}

func TestReadFileLinesReturns10ElementArray(t *testing.T) {
	have, err := readFileLines("data.txt")
	want := 10

	if err != nil {
		t.Errorf("have %#q", err)
	}

	if (len(have) != want) {
		t.Errorf("have %d want %d", len(have), want)
	}
}

func TestGetFieldLengthZeroIndex(t *testing.T) {
	have,err := getFieldLength("1-10", true)
	want := FieldLength{0, 10}

	if err != nil {
		t.Errorf("have %#q", err)
	}

	if have.start != want.start {
		t.Errorf("have %d want %d", have.start, want.start)
	}

	if have.stop != want.stop {
		t.Errorf("have %d want %d", have.stop, want.stop)
	}
}

func TestGetFieldLength(t *testing.T) {
	have, err := getFieldLength("1-10", false)
	want := FieldLength{1, 11}

	if err != nil {
		t.Errorf("have an err %#q", err)
	}

	if have.start != want.start {
		t.Errorf("have %d want %d", have.start, want.start)
	}

	if have.stop != want.stop {
		t.Errorf("have %d want %d", have.stop, want.stop)
	}
}

func TestFieldLengthMapReturnsMap(t *testing.T) {
	have, err := getFieldLengthMap("abc", "1-10", false)
	want := map[string]FieldLength{"abc": {1, 11}}

	if err != nil {
		t.Errorf("have an err %#q", err)
	}

	if have == nil {
		t.Errorf("have nil map")
	}

	if want["abc"] != have["abc"] {
		t.Errorf("have %d want %d", have["abc"], want["abc"])
	}
}

func TestFieldLengthMapReturnsMapZeroIndex(t *testing.T) {
	have, err := getFieldLengthMap("abc", "1-10", true)
	want := map[string]FieldLength{"abc": {0, 10}}

	if err != nil {
		t.Errorf("have an err %#q", err)
	}

	if have == nil {
		t.Errorf("have nil map")
	}

	if want["abc"] != have["abc"] {
		t.Errorf("have %d want %d", have["abc"], want["abc"])
	}
}
