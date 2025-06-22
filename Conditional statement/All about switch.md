# 🧭 All About `switch` 

Go’s `switch` is a multi-way branch statement that evaluates an expression and executes the case that matches.

---

## ✅ 1. Basic `switch` Syntax

```go
switch variable {
case value1:
    // code block
case value2:
    // another block
default:
    // fallback
}
```

### Example:

```go
day := "Tuesday"

switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("Almost weekend")
default:
    fmt.Println("It's", day)
}
```

---

## ✅ 2. `switch` Without Expression (like `if-else if` chain)

```go
score := 85

switch {
case score >= 90:
    fmt.Println("Grade A")
case score >= 80:
    fmt.Println("Grade B")
case score >= 70:
    fmt.Println("Grade C")
default:
    fmt.Println("Fail")
}
```

> This works like `if-else if`, checking each condition top to bottom.

---

## ✅ 3. Multiple Values in One `case`

```go
letter := "A"

switch letter {
case "A", "E", "I", "O", "U":
    fmt.Println("Vowel")
default:
    fmt.Println("Consonant")
}
```

> You can group multiple match values with commas.

---

## ✅ 4. `switch` with Short Statement

Just like `if`, you can declare a variable inside `switch`:

```go
switch age := 25; {
case age < 18:
    fmt.Println("Minor")
case age < 65:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}
```

---

## ✅ 5. Fallthrough

Go’s `switch` does **not** fall through by default. You must **explicitly use `fallthrough`**.

```go
num := 1

switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}
```

> Output:

```
One
Two
Three
```

> ⚠️ Use `fallthrough` sparingly — it's not often needed.

---

## ✅ 6. `default` Case

The `default` case is optional and only runs when no case matches.

```go
switch lang := "Go"; lang {
case "Python":
    fmt.Println("You picked Python")
case "Rust":
    fmt.Println("You picked Rust")
default:
    fmt.Println("Language not recognized")
}
```

---

## ✅ 7. Type Switch (Advanced)

Use when working with interfaces to detect the underlying type.

```go
var x interface{} = 42

switch v := x.(type) {
case int:
    fmt.Println("It's an int:", v)
case string:
    fmt.Println("It's a string:", v)
default:
    fmt.Println("Unknown type")
}
```

---

