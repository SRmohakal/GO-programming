# 🔁 All About `for` Loops in Go

Go has **one looping construct**: `for` — no `while`, no `do-while`, but `for` can do it all!

---

## ✅ 1. Basic `for` Loop

Like C or JavaScript:

```go
for i := 0; i < 5; i++ {
    fmt.Println("i =", i)
}
```

### Breakdown:

* `i := 0` → initialization
* `i < 5` → condition (loop while true)
* `i++` → post statement (runs after each iteration)

---

## ✅ 2. `for` as a `while` loop

You can omit the init and post expressions to make it behave like a `while` loop:

```go
i := 0
for i < 5 {
    fmt.Println("i =", i)
    i++
}
```

---

## ✅ 3. Infinite Loop (`while(true)`)

You can omit everything and just write:

```go
for {
    fmt.Println("infinite loop")
}
```

➡️ Use `break` to exit it.

---

## ✅ 4. Loop with `range` (used for arrays, slices, maps, strings, etc.)

### A. Loop over a slice:

```go
nums := []int{10, 20, 30}
for index, value := range nums {
    fmt.Printf("Index %d: Value %d\n", index, value)
}
```

### B. Loop over a map:

```go
m := map[string]int{"apple": 2, "banana": 3}
for key, value := range m {
    fmt.Println(key, "=>", value)
}
```

### C. Loop over a string (UTF-8 safe):

```go
for i, ch := range "Shourov" {
    fmt.Printf("Index %d: Char %c\n", i, ch)
}
```

---

## ✅ 5. `break` and `continue`

* `break` → exits the loop immediately
* `continue` → skips the current iteration

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // stop when i is 5
    }
    if i%2 == 0 {
        continue // skip even numbers
    }
    fmt.Println("i =", i)
}
```

---

## ✅ 6. Nested `for` Loops

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("(%d,%d) ", i, j)
    }
    fmt.Println()
}
```

---

## ✅ 7. Labelled Loops (for breaking outer loop)

```go
outer:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i == 2 && j == 2 {
            break outer // break outer loop, not just inner
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
```

---

## 🧪 Loop With `range` and Only Value

You can ignore the index with `_`:

```go
for _, value := range []string{"Go", "Rust", "Python"} {
    fmt.Println(value)
}
```

---

