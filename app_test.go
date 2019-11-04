package main

import (
	"testing"
)

func TestAppName(t *testing.T) {
	expect := 7
	actual := AppName(0, 7)

	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
}
