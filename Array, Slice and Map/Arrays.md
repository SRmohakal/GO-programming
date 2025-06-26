# 🧠 Arrays

An **array** is a **fixed-length** sequence of elements of the **same type**. It’s rarely used directly in production (slices are preferred), but understanding arrays builds the foundation for slices.

---

## 🔹 1. Declare an Array

### Syntax:

```go
var arrayName [size]Type
```

### Examples:

```go
var a [3]int              // array of 3 integers
b := [3]string{"Go", "is", "cool"} // initialized array
c := [...]float64{3.14, 2.71}       // size inferred
```

---

## 🔹 2. Access Elements (Indexing)

### Syntax:

```go
array[index]
```

```go
a := [3]int{10, 20, 30}
fmt.Println(a[0]) // Output: 10
```

> ⚠️ Indexing starts from `0`, and `a[3]` (in 3-element array) causes a **runtime error**.

---

## 🔹 3. Update Array Elements

```go
a[1] = 100
fmt.Println(a) // Output: [10 100 30]
```

---

## 🔹 4. Array Length

```go
len(a) // returns the size (number of elements)
```

---

## 🔹 5. Loop Through an Array

### Using `for` loop:

```go
for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
}
```

### Using `for range`:

```go
for index, value := range a {
    fmt.Println("Index:", index, "Value:", value)
}
```

---

## 🔹 6. Initialize with Zero Values

By default, arrays are filled with **zero values** of the type:

```go
var nums [5]int     // [0 0 0 0 0]
var words [3]string // ["", "", ""]
```

---

## 🔹 7. Copying Arrays (By Value)

Arrays are **value types** in Go.

```go
a := [3]int{1, 2, 3}
b := a // copy created
b[0] = 100
fmt.Println(a) // [1 2 3] – original unaffected
fmt.Println(b) // [100 2 3]
```

---

## 🔹 8. Multidimensional Arrays

```go
var matrix [2][3]int
matrix[0][1] = 5
```

### Example:

```go
table := [2][2]string{
    {"Go", "Lang"},
    {"Fast", "Fun"},
}
fmt.Println(table[1][1]) // Output: Fun
```

---

## 🔹 9. Compare Arrays

Arrays of the same type and length can be compared directly.

```go
a := [2]int{1, 2}
b := [2]int{1, 2}
fmt.Println(a == b) // true
```

---

## 🔹 10. Array as Function Argument

Since arrays are copied, modifying inside a function won’t affect the original.

```go
func modify(arr [3]int) {
    arr[0] = 999
}

a := [3]int{1, 2, 3}
modify(a)
fmt.Println(a[0]) // still 1
```

---

## 🔹 11. Sorting Arrays

To sort arrays, you must first **convert them to slices**:

```go
import "sort"

a := [5]int{4, 1, 5, 2, 3}
slice := a[:]         // convert array to slice
sort.Ints(slice)      // sort slice
fmt.Println(a)        // [1 2 3 4 5]
```

> Arrays are fixed-length, so sort works on slices created from arrays.

---

## 🔹 12. Searching in Arrays

### Example:

```go
func contains(arr [5]int, target int) bool {
    for _, val := range arr {
        if val == target {
            return true
        }
    }
    return false
}
```

---

## 🔹 13. Utility: Sum of Array

```go
func sumArray(arr [5]int) int {
    sum := 0
    for _, val := range arr {
        sum += val
    }
    return sum
}
```

---
### 🔸 Built-in Functions:

```go
append(s, 4)       // adds element, returns new slice
len(s), cap(s)     // length and capacity
```

### 🔹 Example:

```go
s := []int{1, 2}
s = append(s, 3)
fmt.Println(s) // [1 2 3]
```

### 🔹 Using `make()`:

```go
nums := make([]int, 5)      // length 5, zero-filled
names := make([]string, 2)  // slice of strings
```

### 🔹 Copying:

```go
a := []int{1, 2, 3}
b := make([]int, len(a))
copy(b, a)
```

