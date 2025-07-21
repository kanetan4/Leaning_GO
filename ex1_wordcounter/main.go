// Goal:
// Build a CLI tool that:
// Accepts a string input
// Splits the string into words
// Counts how many times each word appears
// Prints the word frequency in sorted order (optional)

package main

import "fmt"
import "os"
import "wordcounter/counter"

func main() {
	var input = os.Args[1]
	ans := counter.WordCounter(input)

	for key, value := range ans {
		fmt.Println(key, " => ", value)
	}

}