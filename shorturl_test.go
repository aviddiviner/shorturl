package shorturl

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestCustomAlphabet(t *testing.T) {
	encoder, err := NewEncoder("ab", DefaultBlockSize)
	if err != nil {
		t.Fatal(err)
	}

	url := encoder.EncodeID(12, MinLength)
	expected := "bbaaaaaaaaaaaaaaaaaaaa"
	if url != expected {
		t.Errorf("Expected %s, got %s", expected, url)
	}

	key := encoder.DecodeID(url)
	if key != 12 {
		t.Errorf("Expected %d, got %d", 12, key)
	}
}

func TestShortAlphabet(t *testing.T) {
	_, err := NewEncoder("a", DefaultBlockSize)
	if err == nil {
		t.Error("Expected error for short alphabet, got none")
	}

	_, err = NewEncoder("aa", DefaultBlockSize)
	if err == nil {
		t.Error("Expected error for duplicate characters in alphabet, got none")
	}
}

func TestCalculatedValues(t *testing.T) {
	file, err := os.Open(filepath.Join("testdata", "key_values.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")

		key, _ := strconv.Atoi(parts[0])
		expectedValue := parts[1]

		encodedValue := EncodeID(key)
		if encodedValue != expectedValue {
			t.Errorf("For key %d expected value %s, got %s", key, expectedValue, encodedValue)
		}

		decodedKey := DecodeID(encodedValue)
		if decodedKey != key {
			t.Errorf("For value %s expected key %d, got %d", encodedValue, key, decodedKey)
		}
	}
}
