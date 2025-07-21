// Goal:
// Build a CLI tool that:
// Accepts a string input
// Splits the string into words
// Counts how many times each word appears
// Prints the word frequency in sorted order (optional)

package main

import "fmt"
import "os"
import "sort"
import "wordcounter/counter"

func main() {
	var input = os.Args[1]
	hashmap := counter.WordCounter(input)

	keys := make([]string, 0)
	for key, _ := range hashmap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, " => ", hashmap[key])
	}

}