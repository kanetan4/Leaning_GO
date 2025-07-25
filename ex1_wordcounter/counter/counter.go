package counter

import "strings"

func WordCounter(input string) map[string]int {
	words := strings.Split(input, " ")

	hashmap := make(map[string]int)

	for i := 0; i < len(words); i++ {
		_, exists := hashmap[words[i]]
		if exists {
			hashmap[words[i]]++
		} else {
			hashmap[words[i]] = 1
		}
	}

	return hashmap
}