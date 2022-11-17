package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestMakeUI(t *testing.T) {
	label, entry := makeUI()
	if label.Text != "Hello World" {
		t.Error("Incorrect label text")
	}

	test.Type(entry, "Aaric")
	if label.Text != "Hello Aaric" {
		t.Error("Incorrect entry input")
	}
}
