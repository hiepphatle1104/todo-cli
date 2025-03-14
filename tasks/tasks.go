package tasks

import (
	"fmt"
	. "todo-cli/models"
	. "todo-cli/storage"
)

func AddTask(taskName string) {
	list, err := LoadList()
	if err != nil {
		fmt.Println(err)
		return
	}
	newTaskId := len(list.Tasks) + 1
	var _newTask = Task{Name: taskName, IsDone: false, TaskId: newTaskId}
	list.Tasks = append(list.Tasks, _newTask)

	saveErr := SaveToFile(list)
	if saveErr != nil {
		fmt.Println(saveErr)
		return
	}
	fmt.Println("Added!")
}

func DeleteTask(taskId int) {
	list, err := LoadList()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, task := range list.Tasks {
		if task.TaskId == taskId {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			err = SaveToFile(list)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Deleted!")
			return
		}
	}

	fmt.Println("Task not found")
}

func MarkDone(taskId int) {
	list, err := LoadList()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, task := range list.Tasks {
		if task.TaskId == taskId {

			list.Tasks[i].IsDone = true
			err = SaveToFile(list)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Marked done!")
			return
		}
	}
	fmt.Println("Task not found")
}

func DisplayList() {
	// Load list
	list, err := LoadList()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(list.Tasks) == 0 {
		fmt.Println("List is empty")
		return
	}
	for _, task := range list.Tasks {
		fmt.Printf("%d. %s - Done: %v\n", task.TaskId, task.Name, task.IsDone)
	}
}
