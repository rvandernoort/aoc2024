package framework

import (
	"testing"
)

func TestReadData(t *testing.T) {
	var lines []string
	t.Run("check no panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("ReadData() panicked: %v", r)
			}
		}()

		lines = ReadData("test_input.txt")
	})

	// Your existing test
	if lines[0] != "test1" {
		t.Errorf("Expected 'test1', got '%s'", lines[0])
	}
	if lines[1] != "test2 test3" {
		t.Errorf("Expected 'test2 test3', got '%s'", lines[1])
	}
	if lines[2] != "test4" {
		t.Errorf("Expected 'test4', got '%s'", lines[2])
	}
	if lines[3] != "" {
		t.Errorf("Expected '', got '%s'", lines[3])
	}
	if lines[4] != "test5" {
		t.Errorf("Expected 'test5', got '%s'", lines[4])
	}
}

func TestReadDataEmptyFilename(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReadData() did not panic")
		}
	}()

	ReadData("")
}

func TestReadDataMissingFile(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReadData() did not panic")
		}
	}()

	ReadData("nonexistent.txt")
}
