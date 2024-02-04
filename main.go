package main

import (
	"fmt"
	features "gambit/Features"
	"os"
)

func displayChoices() {
	fmt.Println("Select one of the choices")
	fmt.Println("1. Hello")
	fmt.Println("2. Insert")
	fmt.Println("3. Cost of all project")
	fmt.Println("4. List")
	fmt.Println("5. Update")
	fmt.Println("6. Delete")
	fmt.Println("7. Exit")
}

func main() {
	fmt.Println("Welcome to student DB")
	class := features.NewClass()
	for {

		displayChoices()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:

			fmt.Println(class.Hello())

		case 2:

			var project_name, assigned_to string
			var status bool
			var cost int

			fmt.Println("Project Name (ex: portfolio):")
			fmt.Scanln(&project_name)

			fmt.Println("Assigned To (ex: Sam):")
			fmt.Scanln(&assigned_to)

			fmt.Println("Completion Status (ex: true/false):")
			fmt.Scanln(&status)

			fmt.Println("Cost (ex: 10):")
			fmt.Scanln(&cost)

			class.Insert(cost, status, assigned_to, project_name)
			fmt.Println()

		case 3:

			class.Sum()
			fmt.Println()

		case 4:

			class.List()
			fmt.Println()

		case 5:
			var project_name, assigned_to string
			var cost, index int
			var completion_status bool

			fmt.Println("Index to be deleted:")
			fmt.Scanln(&index)
			fmt.Println("Project Name:")
			fmt.Scanln(&project_name)
			fmt.Println("Assigned To:")
			fmt.Scanln(&assigned_to)
			fmt.Println("Cost:")
			fmt.Scanln(&cost)
			fmt.Println("Completion Status:")
			fmt.Scanln(&completion_status)

			class.Update(completion_status, assigned_to, project_name, index, cost)
			fmt.Println()

		case 6:
			var index int
			fmt.Println("Enter the index to be deleted:")
			fmt.Scanln(&index)
			class.Delete(index)
			fmt.Println()

		case 7:
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
		}
	}
}
