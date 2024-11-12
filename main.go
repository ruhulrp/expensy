package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var expense Expense
	var choice int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What do you want? 1. ADD New Expense, 2. Update Expense, 3. Delete Expense, 4. Get a Report")
	fmt.Scan(&choice)

	switch {
	case choice == 1:
		fmt.Println("Give an ID: ")
		fmt.Scan(&expense.ID)
		fmt.Println("Give a Date: ")
		fmt.Scan(&expense.Date)
		fmt.Println("Give a Category: ")
		category, _ := reader.ReadString('\n')
		expense.Category = category
		fmt.Println("Give an Item: ")
		item, _ := reader.ReadString('\n')
		expense.Item = item
		fmt.Println("Give a Price: ")
		fmt.Scan(&expense.Price)
		fmt.Println("Give a Description: ")
		descripton, _ := reader.ReadString('\n')
		expense.Description = descripton

		fmt.Print(expense.addExpense())

	case choice == 2:
		fmt.Println("You successfully updated an expense!")
	case choice == 3:
		fmt.Println("You successfully deleted an expense!")
	case choice == 4:

		fmt.Println("Here is your Report!")
		fmt.Println(readExpense("expense.json"))
	}

}
