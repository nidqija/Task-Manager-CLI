package main


import(
"fmt"
)


func main() {
	todos := Todos{}
	todos.Add("Buy Milk")
	todos.Add("Buy Water")
	fmt.Printf("%+v\n\n", todos)

	// Delete the first item
	err := todos.Delete(0)
	if err != nil {
		fmt.Println("Failed to delete:", err)
	}

	fmt.Printf("%+v\n\n", todos)
}





