package main

import (
	"errors"
	"fmt"
	"time"
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
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

// ValidateIndex checks if the given index is valid
func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
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