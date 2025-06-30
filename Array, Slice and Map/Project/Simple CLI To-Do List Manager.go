package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := []string{}
	categories := make(map[string][]string)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- To-Do List Manager ---")
		fmt.Println("1) Add task")
		fmt.Println("2) Remove task")
		fmt.Println("3) List all tasks")
		fmt.Println("4) List tasks by category")
		fmt.Println("5) Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task description: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)

			fmt.Print("Enter category for task (or leave empty): ")
			cat, _ := reader.ReadString('\n')
			cat = strings.TrimSpace(cat)

			if cat != "" {
				categories[cat] = append(categories[cat], task)
			}

			fmt.Println("Task added.")

		case "2":
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
				continue
			}
			fmt.Println("Tasks:")
			for i, t := range tasks {
				fmt.Printf("%d) %s\n", i+1, t)
			}
			fmt.Print("Enter task number to remove: ")
			var num int
			fmt.Scanf("%d\n", &num)
			if num < 1 || num > len(tasks) {
				fmt.Println("Invalid task number.")
				continue
			}
			removed := tasks[num-1]
			// Remove from tasks slice
			tasks = append(tasks[:num-1], tasks[num:]...)

			// Also remove from categories
			for cat, ts := range categories {
				for i, t := range ts {
					if t == removed {
						categories[cat] = append(ts[:i], ts[i+1:]...)
						break
					}
				}
			}

			fmt.Println("Task removed.")

		case "3":
			if len(tasks) == 0 {
				fmt.Println("No tasks.")
				continue
			}
			fmt.Println("All Tasks:")
			for i, t := range tasks {
				fmt.Printf("%d) %s\n", i+1, t)
			}

		case "4":
			if len(categories) == 0 {
				fmt.Println("No categories.")
				continue
			}
			fmt.Println("Categories:")
			for cat := range categories {
				fmt.Println("-", cat)
			}
			fmt.Print("Enter category name to view tasks: ")
			catInput, _ := reader.ReadString('\n')
			catInput = strings.TrimSpace(catInput)
			ts, ok := categories[catInput]
			if !ok || len(ts) == 0 {
				fmt.Println("No tasks in this category.")
				continue
			}
			fmt.Printf("Tasks in %s:\n", catInput)
			for i, t := range ts {
				fmt.Printf("%d) %s\n", i+1, t)
			}

		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
