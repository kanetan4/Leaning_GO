package task

import "fmt"

type Task struct {
	Title string
	Priority int
	Completion bool
}

func NewTask(title string, priority int) Task {
	return Task{
		Title: title,
		Priority: priority,
	}
}

func AddTask(input Task) {
	existingTasks, _ := LoadTasks()
	existingTasks = append(existingTasks, input)
	SaveTask(existingTasks) 
	fmt.Println("Task Added")
}

func ListTasks() []Task {
	existingTasks, _ := LoadTasks()
	return existingTasks
}

func CompleteTask(taskName string) {
	existingTasks, _ := LoadTasks()
	for task := range existingTasks {
		fmt.Println(existingTasks[task].Title)
		if existingTasks[task].Title == taskName {
			existingTasks[task].Completion = true
		}
	}
	SaveTask(existingTasks)
}