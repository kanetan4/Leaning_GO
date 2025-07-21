// Goal:
// Build a command-line tool that lets users:
// Add tasks with a name and priority
// List all tasks
// Mark tasks as complete
// Save tasks to a file
// Load tasks from a file

package main

import "fmt"
import "os"
import "main/task"
import "reflect"
import "strconv"

func main() {
	input := os.Args[1:]
	// task.LoadTasks()

	if len(input) == 0 {
		fmt.Println("Please select from the options (add, list, complete)")
	} else {
		switch input[0] {
		case "add":
			fmt.Println("Add")
			if len(input) == 3 {
				if reflect.ValueOf(input[1]).Kind() == reflect.String {
					priority, err := strconv.Atoi(input[2])
					if err != nil {
						fmt.Println("Please input a number instead!!")
					} else {
						newTask := task.NewTask(input[1], priority)
						task.AddTask(newTask)
					}
				}
			}

		case "list":
			fmt.Println("List")
			fmt.Println(task.ListTasks())

		case "complete":
			fmt.Println("Complete")
			if len(input) == 2 {
				if reflect.ValueOf(input[1]).Kind() == reflect.String {
					task.CompleteTask(input[1])
				}
			}

		default:
			fmt.Println("Please select from the options (add, list, complete)")
		}
	}

}