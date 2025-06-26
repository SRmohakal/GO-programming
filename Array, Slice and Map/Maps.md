# 🧠 Maps

---

## 📦 What is a `map`?

A **map** is an **unordered collection** of key-value pairs, where:

* Keys must be **comparable** (e.g. strings, ints, bools).
* Values can be of **any type**.

---

## 📌 1. Declaring and Initializing Maps

### 🔹 Using `make`:

```go
m := make(map[string]int)
```

### 🔹 Using map literal:

```go
studentAges := map[string]int{
    "Alice": 20,
    "Bob":   22,
}
```

---

## 📌 2. Inserting and Updating Elements

```go
m := make(map[string]string)
m["name"] = "Shourov"         // insert
m["name"] = "Shourov Roy"     // update
```

---

## 📌 3. Accessing Elements

```go
value := m["name"]
fmt.Println(value) // Output: Shourov Roy
```

If key doesn’t exist, returns **zero value** of the value type.

---

## 📌 4. Checking if Key Exists (2-value form)

```go
value, exists := m["age"]
if exists {
    fmt.Println("Found:", value)
} else {
    fmt.Println("Not found")
}
```

> Important to avoid confusion between a missing key and a key with zero value (like `0`, `""`, or `false`).

---

## 📌 5. Deleting Elements

```go
delete(m, "name")
```

No error if the key doesn’t exist — it silently does nothing.

---

## 📌 6. Getting Map Length

```go
fmt.Println(len(m)) // number of key-value pairs
```

---

## 📌 7. Iterating Over a Map

```go
for key, value := range studentAges {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}
```

> Note: **Map iteration order is random** and **not guaranteed**.

---

## 📌 8. Nested Maps

You can create maps of maps:

```go
users := map[string]map[string]string{
    "john": {"age": "30", "city": "Dhaka"},
    "jane": {"age": "28", "city": "Sylhet"},
}
fmt.Println(users["john"]["city"]) // Output: Dhaka
```

---

## 📌 9. Map with Struct Values

```go
type Student struct {
    Name string
    Age  int
}

m := map[string]Student{
    "id1": {"Shourov", 22},
    "id2": {"Roy", 23},
}
```

---

## 📌 10. Maps are Reference Types

When passed to functions, changes inside the function **affect the original map**.

```go
func modify(m map[string]int) {
    m["new"] = 100
}

data := map[string]int{}
modify(data)
fmt.Println(data["new"]) // 100
```

---

## 📌 11. Zero Value of a Map

The zero value of a map is `nil`. A `nil` map behaves like an empty map, **except you cannot add to it**.

```go
var m map[string]int     // nil map
// m["a"] = 1            // ❌ panic: assignment to entry in nil map
```

Always use `make()` or literals to initialize maps before use.

---

