// store_test.go
package main

import (
	"testing"
	notes "ex4_noteservice/kitex_gen/notes"
)

func TestSaveAndLoadNotes(t *testing.T) {
	testNotes := []*notes.Note{
		{Id: 1, User: "kane", Title: "Test Note", Content: "This is a test"},
	}

	err := SaveNotes(testNotes)
	if err != nil {
		t.Fatalf("Failed to save notes: %v", err)
	}

	loadedNotes, err := LoadNotes()
	if err != nil {
		t.Fatalf("Failed to load notes: %v", err)
	}

	if len(loadedNotes) != 1 {
		t.Errorf("Expected 1 note, got %d", len(loadedNotes))
	}
}
