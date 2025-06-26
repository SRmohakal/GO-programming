# 🧠 Functions and Return Values

---

## ✅ 1. Defining a Function

### 🔹 Basic Syntax:

```go
func functionName(parameters) returnType {
    // function body
}
```

---

### 🔹 Example:

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

```go
greet("Shourov") // Output: Hello, Shourov
```

---

## ✅ 2. Function with Return Value

```go
func square(n int) int {
    return n * n
}
```

```go
result := square(5)
fmt.Println("Square is", result) // Output: Square is 25
```

---

## ✅ 3. Multiple Return Values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

---

## ✅ 4. Named Return Values

```go
package main

import "fmt"

func stats(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return
}
func main() {
	var sum, product = stats(2, 3)
	fmt.Println("The sum is: ", sum)
	fmt.Println("The product is: ", product)
}
```

---

## ✅ 5. Variadic Functions (Multiple Args)

```go

package main

import "fmt"

func SumAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
func main() {
	var sum = SumAll(1, 2, 3, 4, 5)
	fmt.Println("The sum is: ", sum)
}
```


---

## ✅ 6. Functions as Values

```go
func add(a, b int) int {
    return a + b
}

var operation func(int, int) int = add

fmt.Println(operation(3, 4)) // Output: 7
```

---

## ✅ 7. Anonymous Functions

```go
square := func(n int) int {
    return n * n
}
fmt.Println(square(6)) // Output: 36
```

---

## ✅ 8. Deferred Functions

`defer` delays execution until the surrounding function returns.

```go
func sayHello() {
    defer fmt.Println("Goodbye!")
    fmt.Println("Hello!")
}
```

```
Output:
Hello!
Goodbye!
```

---

## ✅ PRACTICE CHALLENGE

Write a function that takes an integer and returns:

* Its square
* Whether it's even or odd (as a `string`)

```go
func squareAndParity(n int) (int, string) {
    parity := "odd"
    if n%2 == 0 {
        parity = "even"
    }
    return n * n, parity
}
```

```go
sq, parity := squareAndParity(7)
fmt.Println("Square:", sq, "| Parity:", parity)
```

---


