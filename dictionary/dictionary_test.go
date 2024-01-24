package dictionary_test

import (
	"goProject/dictionary"
	"testing"
)

func TestAdd(t *testing.T) {
	d := dictionary.New("test_dictionary.txt")
	defer d.Close()

	// Test success scenario
	err := d.Add("test_word", "test_definition")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test failure scenario - invalid word length
	err = d.Add("", "test_definition")
	if err == nil {
		t.Error("Expected error for invalid word length, but got nil")
	}

	// Test failure scenario - invalid definition length
	err = d.Add("test_word", "")
	if err == nil {
		t.Error("Expected error for invalid definition length, but got nil")
	}
}

func TestGet(t *testing.T) {
	d := dictionary.New("test_dictionary.txt")
	defer d.Close()

	// Add a test entry
	err := d.Add("test_word", "test_definition")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test success scenario
	entry, err := d.Get("test_word")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if entry.Definition != "test_definition" {
		t.Errorf("Expected definition 'test_definition', but got '%s'", entry.Definition)
	}

	// Test failure scenario - word not found
	_, err = d.Get("nonexistent_word")
	if err == nil {
		t.Error("Expected error for nonexistent word, but got nil")
	}
}

func TestRemove(t *testing.T) {
	d := dictionary.New("test_dictionary.txt")
	defer d.Close()

	// Add a test entry
	err := d.Add("test_word", "test_definition")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test success scenario
	err = d.Remove("test_word")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test failure scenario - word not found
	err = d.Remove("nonexistent_word")
	if err == nil {
		t.Error("Expected error for nonexistent word, but got nil")
	}
}
