package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func userGuess(reader *bufio.Reader) int {
	fmt.Print("Enter your number(1-100): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	return num
}

func comp(num, target int) string {
	if num < target {
		return "Too low!"
	} else if num > target {
		return "Too high!"
	} else {
		return "Correct!"
	}
}

func playGame() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	maxAttempts := 5

	fmt.Println("Welcome to the Game!")

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		num := userGuess(reader)
		ans := comp(num, target)
		fmt.Println(ans)

		if ans == "Correct!" {
			fmt.Printf("🏆 You guessed it in %d attempts!\n", attempts)
			return
		} else if attempts < maxAttempts {
			fmt.Printf("⏳ You have %d tries left.\n", maxAttempts-attempts)
		} else {
			fmt.Printf("💀 Out of attempts! The number was %d.\n", target)
		}
	}
}

func main() {
	playGame()
}
