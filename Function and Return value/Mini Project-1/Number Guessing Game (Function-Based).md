# 🧩 Number Guessing Game

### 🎯 Goal:
- The computer generates a random number.
- The user has limited chances to guess it.
- Each guess gives feedback: too high, too low, or correct.
- All logic is organized into **reusable functions**.

---

## ✅ Final Features:
- Use of `rand` for random number
- Modular functions for:
  - Getting user input
  - Checking the guess
  - Running the game loop
- Return values for decision handling

---

## 🔎 Concepts Practiced

| Feature         | Purpose                        |
|----------------|---------------------------------|
| `rand.Intn()`   | Random number generation        |
| `bufio.Reader`  | Full line user input            |
| `func`          | Modular, testable code          |
| `return`        | Used for decision results       |
| `strings.TrimSpace`, `strconv.Atoi` | Clean and convert input |

---

## ✅ Example Run

```
Welcome to the Guessing Game!
Enter your guess (1-100): 50
Too high!
⏳ You have 4 tries left.
Enter your guess (1-100): 25
Too low!
⏳ You have 3 tries left.
Enter your guess (1-100): 37
Correct!
🏆 You guessed it in 3 attempts!
```

---

