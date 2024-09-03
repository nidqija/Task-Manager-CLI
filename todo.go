package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// Todo represents a single task
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// Todos represents a collection of Todo items
type Todos []Todo

// Add adds a new Todo to the list
func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

// ValidateIndex checks if the given index is valid
func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("Invalid Index")
	}
	return nil
}

// Delete removes a Todo by index from the list
func (todos *Todos) Delete(index int) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

// ToggleCompletion toggles the completion status of a Todo item by index
func (todos *Todos) ToggleCompletion(index int) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}

	todo := &(*todos)[index]
	if !todo.Completed {
		completionTime := time.Now()
		todo.CompletedAt = &completionTime
	} else {
		todo.CompletedAt = nil
	}
	todo.Completed = !todo.Completed

	return nil
}

// Edit updates the title of a Todo by index
func (todos *Todos) Edit(index int, title string) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = title
	return nil
}

// Print displays the list of Todos in a table format
func (todos *Todos) Print() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, todo := range *todos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		tbl.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	tbl.Render()
}