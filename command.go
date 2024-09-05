package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CMDFlags struct {
	Add    string
	Edit   string
	Del    int
	Toggle int
	List   bool
}

func NewCmdFlags() *CMDFlags {
	cf := CMDFlags{}

	// Parse command-line flags
	flag.StringVar(&cf.Add, "add", "", "Add a new todo: specify task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title (format: id:new_title)")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle completion")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CMDFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()

	case cf.Add != "":
		todos.Add(cf.Add)
		fmt.Println("Todo added successfully.")
		todos.Print()

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use 'id:new_title'.")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit.")
			os.Exit(1)
		}

		if err := todos.Edit(index, parts[1]); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Todo edited successfully.")
		todos.Print()

	case cf.Toggle != -1:
		if err := todos.ToggleCompletion(cf.Toggle); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Todo toggled successfully.")
		todos.Print()

	case cf.Del != -1:
		if err := todos.Delete(cf.Del); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Todo deleted successfully.")
		todos.Print()

	default:
		fmt.Println("Invalid command.")
		flag.Usage()
		os.Exit(1)
	}
}