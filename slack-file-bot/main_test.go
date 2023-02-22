package main

import "testing"

func TestGetSourcePath(t *testing.T) {
	expectedPath := "/Users/satyapraneel/www/GO/go-projects/slack-file-bot"
	actualPath := getSourcePath()

	if actualPath != expectedPath {
		t.Fatalf("Expected path \"%s\" but got \"%s\"", expectedPath, actualPath)
	}
}
