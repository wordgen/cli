package main

import (
	"flag"
	"testing"
)

func TestCustomFlags(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	wordCase := fs.String("c", "lower", "")
	wordCount := fs.Int("w", 1, "")
	wordSeparator := fs.String("s", " ", "")
	printWithoutNewline := fs.Bool("n", false, "")
	printVersion := fs.Bool("v", false, "")

	err := fs.Parse([]string{"-w", "3", "-c", "title", "-s", "-", "-n"})
	if err != nil {
		t.Fatalf("Error parsing flags: %v", err)
	}

	if *wordCount != 3 {
		t.Errorf("Expected count 3, got %d", *wordCount)
	}
	if *wordCase != "title" {
		t.Errorf("Expected casing title, got %s", *wordCase)
	}
	if *wordSeparator != "-" {
		t.Errorf("Expected separator '-', got %s", *wordSeparator)
	}
	if !*printWithoutNewline {
		t.Errorf("Expected no newline true, got %v", *printWithoutNewline)
	}
	if *printVersion {
		t.Errorf("Expected version flag false, got %v", *printVersion)
	}
}
