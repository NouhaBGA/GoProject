package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"goProject/dictionary"
)

const filePath = "dictionary.txt"

func main() {
	d := dictionary.New(filePath)
	defer d.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter action:")
		action, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		action = strings.TrimSpace(action)

		switch action {
		case "add":
			var wg sync.WaitGroup
			wg.Add(1)
			go actionAdd(d, reader, &wg)
			wg.Wait()
		case "remove":
			var wg sync.WaitGroup
			wg.Add(1)
			go actionRemove(d, reader, &wg)
			wg.Wait()
		case "define":
			actionDefine(d, reader)
		case "list":
			actionList(d)
		case "exit":
			return
		default:
			fmt.Println("Unknown action:", action)
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader, wg *sync.WaitGroup) {
	fmt.Println("Enter word:")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		wg.Done()
		return
	}
	word = strings.TrimSpace(word)

	fmt.Println("Enter definition:")
	definition, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		wg.Done()
		return
	}
	definition = strings.TrimSpace(definition)

	err = d.Add(word, definition)
	if err != nil {
		fmt.Println("Error adding word:", err)
		wg.Done()
		return
	}

	fmt.Printf("Word '%s' added to the dictionary.\n", word)

	// Mark the wait group as done after the asynchronous operation is complete.
	wg.Done()
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Enter word:")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = strings.TrimSpace(word)

	err = d.Remove(word)
	if err != nil {
		fmt.Println("Error removing word:", err)
		return
	}

	fmt.Printf("Word '%s' removed from the dictionary.\n", word)
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Println("Enter word:")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	word = strings.TrimSpace(word)

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Definition of '%s': %s\n", word, entry.String())
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	if err != nil {
		fmt.Println("Error listing words:", err)
		return
	}

	fmt.Println("Words in the dictionary:")
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, entries[word].String())
	}
}
