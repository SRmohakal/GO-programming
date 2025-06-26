# 🧠 Slices

---

## 🔷 What is a Slice?

A **slice** is a **dynamically-sized**, flexible view into an underlying **array**.
It supports:

* Indexing
* Appending
* Resizing
* Copying
* Slicing (like Python or JS)

---

## 📌 1. Declaring a Slice

### 🔹 Slice Literal (most common):

```go
numbers := []int{10, 20, 30}
```

### 🔹 Using `make()`:

```go
s := make([]int, 5)       // length 5, capacity 5
t := make([]int, 3, 10)   // length 3, capacity 10
```

### 🔹 Slicing an Array:

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4] // [2 3 4]
```

---

## 📌 2. Accessing and Modifying Elements

```go
s := []string{"Go", "is", "fun"}
fmt.Println(s[0])  // Go
s[2] = "awesome"   // modify
```

---

## 📌 3. Length and Capacity

```go
fmt.Println(len(s)) // number of elements
fmt.Println(cap(s)) // how many elements underlying array can hold
```

> A slice grows automatically when appended beyond its current length (more on that next).

---

## 📌 4. Appending to a Slice

```go
s := []int{1, 2}
s = append(s, 3)            // [1 2 3]
s = append(s, 4, 5, 6)      // [1 2 3 4 5 6]

t := []int{7, 8}
s = append(s, t...)         // spread operator
```

---

## 📌 5. Copying Slices

```go
a := []int{1, 2, 3}
b := make([]int, len(a))
copy(b, a)
```

> `copy()` copies elements from source to destination slice.

---

## 📌 6. Slice Expressions

### Syntax:

```go
s[start:end]     // from start (inclusive) to end (exclusive)
```

### Examples:

```go
a := []int{0, 1, 2, 3, 4, 5}
fmt.Println(a[1:4])   // [1 2 3]
fmt.Println(a[:3])    // [0 1 2]
fmt.Println(a[3:])    // [3 4 5]
```

---

## 📌 7. Changing a Slice Affects the Underlying Array

```go
arr := [4]int{10, 20, 30, 40}
s := arr[1:3]    // [20 30]
s[0] = 99        // arr now: [10 99 30 40]
```

---

## 📌 8. Comparing Slices (Manual)

You **cannot** compare slices directly (`s1 == s2` is not allowed, except with `nil`). Use a loop or `reflect.DeepEqual()`.

```go
import "reflect"

equal := reflect.DeepEqual(s1, s2)
```

---

## 📌 9. Removing an Element from Slice

There’s no built-in remove, but you can do:

```go
s := []int{1, 2, 3, 4, 5}
i := 2 // remove index 2 (value 3)
s = append(s[:i], s[i+1:]...)
fmt.Println(s) // [1 2 4 5]
```

---

## 📌 10. Inserting an Element

```go
s := []int{1, 2, 4, 5}
i := 2
val := 3
s = append(s[:i], append([]int{val}, s[i:]...)...)
fmt.Println(s) // [1 2 3 4 5]
```

---

## 📌 11. Slices of Structs

```go
type Student struct {
    Name string
    Age  int
}

students := []Student{
    {"Shourov", 22},
    {"Roy", 23},
}
```

---

## 📌 12. Iterating Over a Slice

```go
nums := []int{1, 2, 3}
for i, val := range nums {
    fmt.Println(i, val)
}
```

---

## 📌 13. Nil vs Empty Slice

```go
var a []int          // nil slice, len=0, cap=0
b := []int{}         // empty slice, len=0, cap=0
fmt.Println(a == nil) // true
fmt.Println(b == nil) // false
```

---

