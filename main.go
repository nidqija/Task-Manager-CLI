package main


import(
"fmt"

)


func main() {
	// Example usage of the Todos list
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.Add("Buy Milk")
	todos.Add("Complete Golang Project")
	todos.Print()
	fmt.Println()
	storage.Save(todos)
}





