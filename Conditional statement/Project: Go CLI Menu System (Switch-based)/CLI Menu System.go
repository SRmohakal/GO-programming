package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var tasks []string

func calculator(reader *bufio.Reader) {
	fmt.Print("Enter first number: ")
	aStr, _ := reader.ReadString('\n')
	a, _ := strconv.ParseFloat(strings.TrimSpace(aStr), 64)

	fmt.Print("Enter an operator (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	fmt.Print("Enter second number: ")
	bStr, _ := reader.ReadString('\n')
	b, _ := strconv.ParseFloat(strings.TrimSpace(bStr), 64)

	switch op {
	case "+":
		fmt.Println("Result:", a+b)
	case "-":
		fmt.Println("Result:", a-b)
	case "*":
		fmt.Println("Result:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("❌ Cannot divide by zero")
		}
	default:
		fmt.Println("❌ Invalid operator")
	}
}

func taskManager(reader *bufio.Reader) {
	for {
		fmt.Println("\n📋 Task Manager:")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Clear Tasks")
		fmt.Println("4. Back to Main Menu")
		fmt.Print("Choose: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, strings.TrimSpace(task))
			fmt.Println("✅ Task added!")
		case "2":
			fmt.Println("📝 Your Tasks:")
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				for i, t := range tasks {
					fmt.Printf("%d. %s\n", i+1, t)
				}
			}
		case "3":
			tasks = []string{}
			fmt.Println("🗑️ All tasks cleared!")
		case "4":
			return
		default:
			fmt.Println("❌ Invalid option")
		}
	}
}

func showClock() {
	now := time.Now()
	fmt.Println("🕒 Current Time:", now.Format("15:04:05"))
}

func greetUser(reader *bufio.Reader) {
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("👋 Hello,", name)
}

func showGoInfo() {
	fmt.Println("🐹 Go is simple, fast, compiled, and has great concurrency support.")
}

func tellJoke() {
	fmt.Println("😂 Debugging: Being the detective in a crime movie where you're also the murderer.")
}

func main() {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n📋 Main Menu:")
		fmt.Println("1. Greet Me")
		fmt.Println("2. Show Go Info")
		fmt.Println("3. Tell a Joke")
		fmt.Println("4. Calculator")
		fmt.Println("5. Task Manager")
		fmt.Println("6. Show Current Time")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := input.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			greetUser(input)
		case "2":
			showGoInfo()
		case "3":
			tellJoke()
		case "4":
			calculator(input)
		case "5":
			taskManager(input)
		case "6":
			showClock()
		case "7":
			fmt.Println("👋 Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice. Try again.")
		}
	}
}
