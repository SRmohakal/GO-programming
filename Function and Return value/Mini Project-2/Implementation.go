package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func logToFile(entry string) {
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Logging error: ", err)
		return
	}
	defer f.Close()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	f.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, entry))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func main() {
	for {
		fmt.Println("\n=== Mini Utility Toolbox ===")
		fmt.Println("📐 Math Utilities")
		fmt.Println("  1. Prime Check")
		fmt.Println("  2. Factorial")
		fmt.Println("🔤 Text Utilities")
		fmt.Println("  3. Palindrome Check")
		fmt.Println("🌡️  Temperature Utilities")
		fmt.Println("  4. Celsius to Fahrenheit")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var num int
			fmt.Print("Enter a number: ")
			fmt.Scanln(&num)
			result := isPrime(num)
			log := fmt.Sprintf("Prime Check (%d): %v", num, result)
			logToFile(log)
			if result {
				fmt.Println(num, "is a Prime Number")
			} else {
				fmt.Println(num, "is NOT a Prime Number")
			}

		case 2:
			var num int
			fmt.Print("Enter a number: ")
			fmt.Scanln(&num)
			result := factorial(num)
			log := fmt.Sprintf("Factorial (%d): %d", num, result)
			logToFile(log)
			fmt.Println("Factorial:", result)

		case 3:
			var word string
			fmt.Print("Enter a word: ")
			fmt.Scanln(&word)
			result := isPalindrome(word)
			log := fmt.Sprintf("Palindrome Check (%s): %v", word, result)
			logToFile(log)
			if result {
				fmt.Println(word, "is a Palindrome")
			} else {
				fmt.Println(word, "is NOT a Palindrome")
			}

		case 4:
			var celsius float64
			fmt.Print("Enter Celsius temperature: ")
			fmt.Scanln(&celsius)
			fahrenheit := celsiusToFahrenheit(celsius)
			log := fmt.Sprintf("Celsius to Fahrenheit (%.2f°C): %.2f°F", celsius, fahrenheit)
			logToFile(log)
			fmt.Printf("Fahrenheit: %.2f°F\n", fahrenheit)

		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
