# ✅ Conditional Statements: `if`, `else if`, `else`

Go provides powerful and clean control flow using these:

---

## 🔹 1. `if` Statement

```go
if condition {
    // execute if condition is true
}
```

### Example:

```go
x := 10
if x > 5 {
    fmt.Println("x is greater than 5")
}
```

---

## 🔹 2. `if` with a short variable declaration

You can declare and use a variable *only inside* the `if` block:

```go
if y := 20; y > 10 {
    fmt.Println("y is greater than 10")
}
```

> The variable `y` is only available inside that `if`/`else if`/`else` chain.

---

## 🔹 3. `if-else`

```go
if condition {
    // true block
} else {
    // false block
}
```

### Example:

```go
score := 40
if score >= 50 {
    fmt.Println("Passed")
} else {
    fmt.Println("Failed")
}
```

---

## 🔹 4. `if-else if-else`

```go
if condition1 {
    // block1
} else if condition2 {
    // block2
} else {
    // fallback block
}
```

### Example:

```go
grade := 85
if grade >= 90 {
    fmt.Println("Grade: A")
} else if grade >= 80 {
    fmt.Println("Grade: B")
} else if grade >= 70 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}
```

---

## 🔹 5. Nesting `if` statements

```go
x := 5
y := 10
if x < 10 {
    if y > 5 {
        fmt.Println("x < 10 and y > 5")
    }
}
```

> ⚠️ Too much nesting can make code hard to read — prefer `else if`.

---


## 🔹 6. Example: Age Category Checker

```go
age := 25

if age < 13 {
    fmt.Println("Child")
} else if age < 20 {
    fmt.Println("Teen")
} else if age < 60 {
    fmt.Println("Adult")
} else {
    fmt.Println("Senior")
}
```

---

## 🛡 Best Practices

✅ Always use `{}` even for one-liners — it improves readability
✅ Use short declarations (`if x := ...; condition`) when useful
✅ Don’t over-nest; use `else if` or early returns
✅ Avoid deeply nested `if` chains — split into functions if complex

---

