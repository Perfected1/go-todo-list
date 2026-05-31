package main

import (
	"fmt"
	"os"

	"cli-todo/todo"
)

const fileName = "todos.json"

func main() {
	// Load existing todos from file at startup
	todos, err := todo.LoadTodos(fileName)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	// Basic usage guide when no command is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add <task>       - Add a new todo")
		fmt.Println("  list             - Show all todos")
		fmt.Println("  complete <id>    - Mark todo as completed")
		fmt.Println("  delete <id>      - Delete a todo")
		return
	}

	command := os.Args[1]

	switch command {

	// Add a new todo
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task.")
			return
		}

		task := os.Args[2]

		todos.Add(task)

		if err := todos.Save(fileName); err != nil {
			fmt.Println("Error saving todo:", err)
			return
		}

		fmt.Println("Todo added successfully.")

	// List all todos
	case "list":
		todos.List()

	// Mark todo as completed
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide todo ID.")
			return
		}

		var todoID int
		_, err := fmt.Sscanf(os.Args[2], "%d", &todoID)
		if err != nil {
			fmt.Println("Invalid ID format.")
			return
		}

		ok := todos.Complete(todoID)
		if !ok {
			fmt.Println("Todo not found.")
			return
		}

		if err := todos.Save(fileName); err != nil {
			fmt.Println("Error saving file:", err)
			return
		}

		fmt.Println("Todo marked as completed.")

	// Delete a todo
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide todo ID.")
			return
		}

		var todoID int
		_, err := fmt.Sscanf(os.Args[2], "%d", &todoID)
		if err != nil {
			fmt.Println("Invalid ID format.")
			return
		}

		ok := todos.Delete(todoID)
		if !ok {
			fmt.Println("Todo not found.")
			return
		}

		if err := todos.Save(fileName); err != nil {
			fmt.Println("Error saving file:", err)
			return
		}

		fmt.Println("Todo deleted successfully.")

	// Unknown command handler
	default:
		fmt.Println("Unknown command.")
	}
}