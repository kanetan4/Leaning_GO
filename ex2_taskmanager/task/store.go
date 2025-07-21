package task

import "encoding/json"
import "os"
// import "fmt"

const fileName = "tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(fileName) // read json
	
	if err != nil {
		if os.IsNotExist(err) { //if file does not exist
			return []Task{}, err //return empty list
		}
	}

	err = json.Unmarshal(data, &tasks) //take data from json save into mem address at tasks
	return tasks, err
}

func SaveTask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ") //basically jsonify for GO (data, prefix, indent)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644) // write in
}