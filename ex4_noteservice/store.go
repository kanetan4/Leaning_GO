package main

import (
	"encoding/json"
	"os"
	"sync"
	notes "ex4_noteservice/kitex_gen/notes"
)

const fileName = "notes.json"

var mutex sync.Mutex

func LoadNotes() ([]*notes.Note, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var list []*notes.Note
	data, err := os.ReadFile(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			return []*notes.Note{}, nil
		}
	}

	if len(data) == 0 {
		return []*notes.Note{}, nil
	}

	err = json.Unmarshal(data, &list) //take data from json save into mem address at tasks
	return list, err
}

func SaveNotes(notes []*notes.Note) error {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := json.MarshalIndent(notes, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}