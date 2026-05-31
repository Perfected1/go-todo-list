# 📝 CLI Todo App (Go)

A simple and lightweight command-line Todo application built with Go.  
It allows you to manage tasks directly from your terminal with persistent storage using a JSON file.

---

## 🚀 Features

- Add new todos
- View all todos
- Mark todos as completed
- Delete todos
- Persistent storage using `todos.json`
- Simple and fast CLI interface

---

## 🛠 Tech Stack

- Go (Golang)
- File handling (JSON)
- CLI (os.Args)

---

## 📁 Project Structure

```text
cli-todo/
├── main.go
├── go.mod
├── todos.json
└── todo/
    └── todo.go
````

---

## ⚙️ Installation

Clone the repository:

```bash id="c8l2bq"
git clone https://github.com/Perfected1/go-todo-list.git
cd go-todo-list
```

Install dependencies:

```bash id="9n5r2k"
go mod tidy
```

---

## ▶️ Usage

Run the application:

```bash id="m8xq1d"
go run .
```

---

## 📌 Commands

### ➕ Add a Todo

```bash id="v3t9pa"
go run . add "Learn Go"
```

---

### 📋 List Todos

```bash id="k2d7ns"
go run . list
```

---

### ✅ Complete a Todo

```bash id="q9p4zc"
go run . complete 1
```

---

### 🗑 Delete a Todo

```bash id="x7m1we"
go run . delete 1
```

---

## 💾 Data Storage

Todos are stored locally in:

```text id="d4l9qv"
todos.json
```

This file is automatically created after the first todo is added.

---

## 🧠 What You Learned

This project demonstrates:

* Go structs and methods
* CLI argument handling
* File I/O operations
* JSON encoding/decoding
* Basic CRUD operations
* Project structuring in Go

---

## 🚀 Future Improvements

* Multi-word input improvement
* Colored terminal output
* Update/edit todos
* Filtering (completed / pending)
* Better CLI parser (flags instead of os.Args)
* Modular architecture (cmd/internal structure)

---

## 👨‍💻 Author

Chike Jerry Nnamadim

say 👍
```
