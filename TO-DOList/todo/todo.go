package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo represents a single task in the system.
// It holds the task data and whether it's completed or not.
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// TodoList is an in-memory collection of todos.
// It also doubles as the structure we persist to JSON.
type TodoList []Todo

// LoadTodos reads todos from a JSON file.
// If the file does not exist, it returns an empty list (fresh start).
func LoadTodos(filename string) (TodoList, error) {
	var todos TodoList

	data, err := os.ReadFile(filename)
	if err != nil {
		// If file doesn't exist yet, return empty list (normal first run behavior)
		if os.IsNotExist(err) {
			return todos, nil
		}
		return nil, err
	}

	// Convert JSON data into TodoList
	err = json.Unmarshal(data, &todos)
	return todos, err
}

// Save writes the current todos into a JSON file.
// This ensures data persists between runs of the program.
func (t *TodoList) Save(filename string) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Add creates a new todo item and appends it to the list.
// ID is generated based on the current maximum ID in the list.
func (t *TodoList) Add(task string) {
	maxID := 0

	// Find the highest existing ID to avoid collisions after deletions
	for _, todo := range *t {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}

	newTodo := Todo{
		ID:        maxID + 1,
		Task:      task,
		Completed: false,
	}

	*t = append(*t, newTodo)
}

// List prints all todos in a clean, readable format.
func (t TodoList) List() {
	if len(t) == 0 {
		println("No todos found yet.")
		return
	}

	for _, todo := range t {
		status := " "
		if todo.Completed {
			status = "✓"
		}

		// Clean CLI output format
		fmt.Printf("[%s] %d - %s\n", status, todo.ID, todo.Task)
	}
}

// Complete marks a todo as completed using its ID.
// Returns true if the todo was found and updated.
func (t *TodoList) Complete(id int) bool {
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Completed = true
			return true
		}
	}
	return false
}

// Delete removes a todo from the list by its ID.
// Returns true if deletion was successful.
func (t *TodoList) Delete(id int) bool {
	for i, todo := range *t {
		if todo.ID == id {

			// Remove item using slice re-build
			*t = append((*t)[:i], (*t)[i+1:]...)
			return true
		}
	}
	return false
}