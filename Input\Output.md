# 🧾 Input & Output (I/O) 

---

## ✅ 1. **Console Input (Standard Input)**

### 🔹 A. Using `fmt.Scan()`

Reads space-separated input values. Stops reading at the first space or newline.

```go
var name string
fmt.Print("Enter name: ")
fmt.Scan(&name)
fmt.Println("You entered:", name)
```

> ❗ Only reads up to first space: `"Shourov Roy"` becomes `"Shourov"`.

---

### 🔹 B. Using `fmt.Scanln()`

Reads until newline. Doesn’t include newline in input.

```go
var name string
fmt.Print("Enter name: ")
fmt.Scanln(&name)
```

> ❗ Still stops at spaces — not ideal for full names or sentences.

---

### 🔹 C. Using `bufio.NewReader` and `ReadString('\n')`

✅ **Best for full-line input**, including spaces.

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter full name: ")
name, _ := reader.ReadString('\n')
name = strings.TrimSpace(name) // Remove \n
fmt.Println("Hello,", name)
```

---

### 🔹 D. Using `fmt.Scanf()` with format specifiers

```go
var name string
fmt.Print("Enter name: ")
fmt.Scanf("%s", &name)
```

> Can format inputs (`%s`, `%d`, etc.), but also stops at space unless using special patterns.

---

## ✅ 2. **Console Output (Standard Output)**

### 🔹 A. `fmt.Print()`, `fmt.Println()`, `fmt.Printf()`

```go
fmt.Print("Hello")      // No newline
fmt.Println("Hello")    // Adds newline
fmt.Printf("Hi %s\n", "Shourov")  // Formatted output
```

---

## ✅ 3. **File Input & Output**

### 🔹 A. Reading a whole file

```go
data, err := os.ReadFile("myfile.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

---

### 🔹 B. Writing to a file (overwrites)

```go
err := os.WriteFile("output.txt", []byte("Hello File!"), 0644)
if err != nil {
    log.Fatal(err)
}
```

---

### 🔹 C. Append to a file

```go
f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer f.Close()

if _, err := f.WriteString("Another line\n"); err != nil {
    log.Fatal(err)
}
```

---

## ✅ 4. **Advanced Input: Reading Numbers or Tokens in Loop**

```go
var n int
fmt.Print("Enter number: ")
fmt.Scanln(&n)

nums := make([]int, n)
for i := 0; i < n; i++ {
    fmt.Scan(&nums[i])
}
fmt.Println("You entered:", nums)
```

---

## ✅ 5. **JSON Input/Output**

### 🔹 A. Marshal (convert struct to JSON)

```go
person := map[string]string{"name": "Shourov", "city": "Dhaka"}
jsonData, _ := json.Marshal(person)
fmt.Println(string(jsonData)) // Output: {"city":"Dhaka","name":"Shourov"}
```

### 🔹 B. Unmarshal (convert JSON to struct)

```go
var person map[string]string
jsonStr := `{"name":"Shourov","city":"Dhaka"}`
json.Unmarshal([]byte(jsonStr), &person)
fmt.Println(person["name"])
```

---

## ✅ 6. **Buffered Writing (Large Output)**

```go
writer := bufio.NewWriter(os.Stdout)
writer.WriteString("Buffered output\n")
writer.Flush()
```

---

