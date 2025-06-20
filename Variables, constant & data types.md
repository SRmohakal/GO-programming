---

## 🧠 **Variables, Constants, and Data Types in Go**

---

### ✅ 1. **Declaring Variables**

Go is statically typed, so every variable has a type.

#### 📌 Using `var`:

```go
package main
import "fmt"

func main() {
    var name string = "Shourov"
    var age int = 21
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

#### 📌 Type Inference (Go guesses the type):

```go
func main() {
    var country = "Bangladesh" // inferred as string
    fmt.Println("Country:", country)
}
```

#### 📌 Short-hand declaration (inside functions only):

```go
func main() {
    city := "Sylhet" // same as var city string = "Sylhet"
    fmt.Println("City:", city)
}
```

---

### ✅ 2. **Constants**

```go
const pi = 3.1416

func main() {
    fmt.Println("PI:", pi)
}
```

---

### ✅ 3. **Data Types**

| Type      | Example                 |
| --------- | ----------------------- |
| `int`     | `var x int = 10`        |
| `float64` | `var pi float64 = 3.14` |
| `string`  | `var s string = "Go"`   |
| `bool`    | `var ok bool = true`    |

---

## 🎯 **Exercise**

Create a file `variables.go` with the following:

```go
package main

import "fmt"

func main() {
    var fullName string = "Shourov Roy"
    age := 22
    const university = "SUST"
    var cgpa float64 = 3.95
    var isGraduated bool = false

    fmt.Println("Name:", fullName)
    fmt.Println("Age:", age)
    fmt.Println("University:", university)
    fmt.Println("CGPA:", cgpa)
    fmt.Println("Graduated?", isGraduated)
}
```

✅ **Run:**

```bash
go run variables.go
```
