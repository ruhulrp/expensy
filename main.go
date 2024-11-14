package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var add_expense Expense
	var choice int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What do you want?\n1. ADD New Expense\n2. Update Expense\n3. Delete Expense\n4. Get a Report")
	fmt.Scan(&choice)

	switch {
	case choice == 1: // ADD New Expense

		fmt.Println("Give an ID: ")
		fmt.Scan(&add_expense.ID)

		fmt.Println("Give a Date: ")
		fmt.Scan(&add_expense.Date)

		fmt.Println("Give a Category: ")
		category, _ := reader.ReadString('\n')
		add_expense.Category = category

		fmt.Println("Give an Item: ")
		item, _ := reader.ReadString('\n')
		add_expense.Item = item

		fmt.Println("Give a Price: ")
		fmt.Scan(&add_expense.Price)

		fmt.Println("Give a Description: ")
		descripton, _ := reader.ReadString('\n')
		add_expense.Description = descripton

		fmt.Print(addExpense(add_expense))

	case choice == 2: // Update an Expense

		var expense_id_to_update int

		fmt.Println("Give an ID: ")
		fmt.Scan(&expense_id_to_update)

		saveUpdatedInfo(expense_id_to_update)

		fmt.Println("You successfully updated an expense!")

	case choice == 3: // Delete an Expense

		var expense_id_to_delete int

		fmt.Println("Give an ID: ")
		fmt.Scan(&expense_id_to_delete)

		deleteExpense(expense_id_to_delete)

		fmt.Println("You successfully deleted an expense!")

	case choice == 4: // Generate a Report

		fmt.Println("Here is your Report:")
		expensesReport()
	}

}
