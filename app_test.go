package main

import "testing"

func TestAppName(t *testing.T) {
	expect := "Learning Golang now!"
	actual := AppName()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
