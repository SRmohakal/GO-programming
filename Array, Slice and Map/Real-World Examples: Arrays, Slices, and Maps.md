# Real-World Examples: Arrays, Slices, and Maps

---

## 1️⃣ Arrays — Example: Days of the Week

Arrays work well when the size is fixed and known in advance.

```go
package main

import "fmt"

func main() {
    // Array of 7 strings representing days of the week
    var days [7]string = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

    fmt.Println("Days of the week:")
    for i := 0; i < len(days); i++ {
        fmt.Printf("%d: %s\n", i+1, days[i])
    }
}
```

*Use case:* Fixed list of constant values, such as days, months, directions, or statuses.

---

## 2️⃣ Slices — Example: Managing a To-Do List

Slices are perfect for lists that can grow/shrink dynamically.

```go
package main

import "fmt"

func main() {
    // Initial empty to-do list
    todos := []string{}

    // Add tasks
    todos = append(todos, "Buy groceries")
    todos = append(todos, "Call Alice")
    todos = append(todos, "Pay bills")

    fmt.Println("To-Do List:")
    for i, task := range todos {
        fmt.Printf("%d. %s\n", i+1, task)
    }

    // Mark task 2 as done (remove "Call Alice")
    indexToRemove := 1
    todos = append(todos[:indexToRemove], todos[indexToRemove+1:]...)

    fmt.Println("\nUpdated To-Do List:")
    for i, task := range todos {
        fmt.Printf("%d. %s\n", i+1, task)
    }
}
```

*Use case:* Any dynamically sized collection like lists, queues, stacks.

---

## 3️⃣ Maps — Example: Inventory Tracking System

Maps work well as dictionaries or key-value stores.

```go
package main

import "fmt"

func main() {
    // Inventory map: product name -> quantity
    inventory := make(map[string]int)

    // Add items
    inventory["Apples"] = 50
    inventory["Bananas"] = 30
    inventory["Oranges"] = 20

    fmt.Println("Inventory Report:")
    for product, qty := range inventory {
        fmt.Printf("%s: %d units\n", product, qty)
    }

    // Update inventory for Apples
    inventory["Apples"] += 15

    // Check if product exists before selling
    productToSell := "Bananas"
    if qty, ok := inventory[productToSell]; ok && qty > 0 {
        inventory[productToSell] -= 5
        fmt.Printf("\nSold 5 %s, remaining %d\n", productToSell, inventory[productToSell])
    } else {
        fmt.Printf("\n%s is out of stock!\n", productToSell)
    }
}
```

*Use case:* Fast lookups for counts, configurations, caching data by keys.

---

##Combining Slices and Maps — User Roles Management

```go
package main

import "fmt"

func main() {
    // Map user IDs to their roles (slice of strings)
    userRoles := make(map[string][]string)

    userRoles["user1"] = []string{"admin", "editor"}
    userRoles["user2"] = []string{"viewer"}
    userRoles["user3"] = []string{"editor", "viewer"}

    // Add a new role to user2
    userRoles["user2"] = append(userRoles["user2"], "contributor")

    // Print user roles
    for user, roles := range userRoles {
        fmt.Printf("%s roles: %v\n", user, roles)
    }
}
```



