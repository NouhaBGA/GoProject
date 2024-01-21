package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return e.Definition
}

type Dictionary struct {
	filePath string
}

func New(filePath string) *Dictionary {
	return &Dictionary{

		filePath: filePath,
	}
}

func (d *Dictionary) Add(word string, definition string) error {
	file, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	entryStr := fmt.Sprintf("%s: %s\n", word, definition)
	_, err = file.WriteString(entryStr)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) Get(word string) (Entry, error) {
	file, err := os.Open(d.filePath)
	if err != nil {
		return Entry{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 && strings.TrimSpace(parts[0]) == word {
			return Entry{Definition: strings.TrimSpace(parts[1])}, nil
		}
	}

	return Entry{}, fmt.Errorf("word not found in the dictionary")
}

func (d *Dictionary) Remove(word string) error {
	file, err := os.Open(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	lines, err := readLines(file)
	if err != nil {
		return err
	}

	for i, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 && strings.TrimSpace(parts[0]) == word {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}
	}

	file, err = os.Create(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Dictionary) List() ([]string, map[string]Entry, error) {
	file, err := os.Open(d.filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	lines, err := readLines(file)
	if err != nil {
		return nil, nil, err
	}

	entries := make(map[string]Entry)
	var words []string

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			word := strings.TrimSpace(parts[0])
			definition := strings.TrimSpace(parts[1])
			entries[word] = Entry{Definition: definition}
			words = append(words, word)
		}
	}

	return words, entries, nil
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
