package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    // Step 1: Get user's name
    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    // Step 2: Get user's note (single line)
    fmt.Print("Enter your note: ")
    note, _ := reader.ReadString('\n')
    note = strings.TrimSpace(note)

    // Step 3: Create file with user's name
    fileName := name + ".txt"
    file, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Step 4: Write note to file
    _, err = file.WriteString(note + "\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("✅ Note saved successfully!")

    // Step 5: Read back the file and display note
    data, err := os.ReadFile(fileName)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println("\n📄 Your Saved Note:")
    fmt.Println(string(data))
}
