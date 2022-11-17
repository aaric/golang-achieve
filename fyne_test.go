package main

import "testing"

func TestMakeUI(t *testing.T) {
	l, _ := makeUI()
	if l.Text != "hello world" {
		t.Error("Incorrect label text")
	}
}
