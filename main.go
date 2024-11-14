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

	fmt.Println("What do you want?\n1. ADD New Expense\n2. Update Expense\n3. Delete Expense\n4. Get a Report")
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

		var n int

		fmt.Println("Give an ID: ")
		fmt.Scan(&n)

		saveUpdatedInfo(n)

		fmt.Println("You successfully updated an expense!")
	case choice == 3:
		var n int

		fmt.Println("Give an ID: ")
		fmt.Scan(&n)

		deleteExpense(n)

		fmt.Println("You successfully deleted an expense!")
	case choice == 4:
		fmt.Println("Here is your Report:")
		expensesReport()
	}

}
