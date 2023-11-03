package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	completed   bool
}

type TaskList struct {
	tasks []Task
}

func (t *TaskList) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskList) markCompleted(index int) {
	t.tasks[index].completed = true
}

func (t *TaskList) updateTask(index int, task Task) {
	t.tasks[index] = task
}

func (t *TaskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {

	taskList := TaskList{}

	for {
		var opt int

		read := bufio.NewReader(os.Stdin)

		fmt.Print("Select an option: \n",
			"1. Add Task\n",
			"2. Update Task\n",
			"3. Delete Task\n",
			"4. Mark Completed\n",
			"5. View Tasks\n",
			"6. Exit\n")

		fmt.Print("Enter the option: ")
		fmt.Scanln((&opt))

		switch opt {

		case 1:
			var task Task
			fmt.Print("Enter the name of the task: ")
			task.name, _ = read.ReadString('\n')

			fmt.Print("Enter the description of the task: ")
			task.description, _ = read.ReadString('\n')

			taskList.AddTask(task)

			fmt.Println("Task added successfully")

		case 2:
			var index int
			var task Task

			fmt.Print("Enter the index of the task to update: ")
			fmt.Scanln(&index)

			fmt.Print("Enter the name of the task: ")
			task.name, _ = read.ReadString('\n')

			fmt.Print("Enter the description of the task: ")
			task.description, _ = read.ReadString('\n')

			taskList.updateTask(index, task)

			fmt.Println("Task updated successfully")

		case 3:
			var index int
			fmt.Print("Enter the index of the task to delete: ")
			fmt.Scanln(&index)

			taskList.deleteTask(index)
			fmt.Println("Task deleted successfully")

		case 4:
			var index int
			fmt.Print("Enter the index of the task to mark as complete: ")
			fmt.Scanln(&index)

			taskList.markCompleted(index)

			fmt.Println("Task marked as completed")

		case 5:

			fmt.Println("Task List: ")
			fmt.Println("==================")

			for i, t := range taskList.tasks {
				fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.name, t.description, t.completed)
			}

			fmt.Println("==================")

		case 6:
			fmt.Println("Bye")
			return

		default:
			fmt.Println("Invalid option")

		}

	}

}
