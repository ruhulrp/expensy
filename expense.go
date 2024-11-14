package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// The information that should be input for every Expense
type Expense struct {
	ID          int
	Date        string
	Category    string
	Item        string
	Price       float32
	Description string
}

// ======================================
// All Data will be saved in this file.
const JSONFILENAME = "expense.json"

var expenses_slice []Expense

//======================================

func addExpense(exp Expense) string {

	// Storing data recieved from user input
	add_new_expense := Expense{ID: exp.ID, Date: exp.Date, Category: exp.Category, Item: exp.Item, Price: exp.Price, Description: exp.Description}

	// Storing expense struct to expenses slice
	expenses_slice = append(expenses_slice, add_new_expense)

	if !fileExist() {
		writeExpense() // Creates fresh JSON file and stores data
	} else {
		appendExpense(add_new_expense) // Updates existing JSON file
	}

	return "You successfully added an expense!\n"

}

func readExpense() {

	file, _ := os.Open(JSONFILENAME)

	defer file.Close()

	decoder := json.NewDecoder(file)

	decoder.Decode(&expenses_slice)

}

func writeExpense() {

	file, _ := os.Create(JSONFILENAME)

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.Encode(expenses_slice)

}

func appendExpense(newExpense Expense) {

	readExpense() // Sets a new value in expense_slice variable

	expenses_slice = append(expenses_slice, newExpense)

	writeExpense() // Save newly added expense with other expenses

}

func deleteExpense(id int) {

	var temp_expense_slice []Expense

	readExpense() // Sets a new value in expense_slice variable

	for expense_index, single_expense := range expenses_slice {

		if single_expense.ID != id {
			temp_expense_slice = append(temp_expense_slice, expenses_slice[expense_index])
		}
	}

	expenses_slice = temp_expense_slice // Setting new value after deleting the target

	writeExpense() // Saves newly created slice

}

func expensesReport() {

	readExpense() // Sets a new value in expense_slice variable

	var data [][]string // Declaring two-dimensional slice of string

	for _, single_expense := range expenses_slice {

		datum := []string{fmt.Sprintf("%d", single_expense.ID), single_expense.Date, single_expense.Category, single_expense.Item, fmt.Sprintf("%.2f", single_expense.Price), single_expense.Description}

		data = append(data, datum)
	}

	report_table_headers := []string{"ID", "Date", "Category", "Item", "Price", "Description"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(report_table_headers)

	for _, row := range data[0:] {

		table.Append(row)
	}

	table.Render()

}

func saveUpdatedInfo(id int) {

	updateExp := getUpdateInfo(id)

	readExpense() // Sets a new value in expense_slice variable

	expenses_slice = append(expenses_slice, updateExp)

	writeExpense() // Saves newly created slice

}

func getUpdateInfo(id int) Expense {

	reader := bufio.NewReader(os.Stdin)

	var choice int

	expense_need_to_update := findExpenseByID(id)

	///////////////////////////////////

	fmt.Println("What do you edit?\n1. Date\n2. Category\n3. Item\n4. Price\n5. Description")
	fmt.Scan(&choice)

	///////////////////////////////

	switch {
	case choice == 1:
		fmt.Println("Edit Date: " + expense_need_to_update.Date)
		fmt.Scan(&expense_need_to_update.Date)

	case choice == 2:
		fmt.Print("Edit Category: " + expense_need_to_update.Category)
		category, _ := reader.ReadString('\n')
		expense_need_to_update.Category = category

	case choice == 3:
		fmt.Print("Edit Item: " + expense_need_to_update.Item)
		item, _ := reader.ReadString('\n')
		expense_need_to_update.Item = item

	case choice == 4:
		fmt.Printf("Edit Price: %v\n", expense_need_to_update.Price)
		fmt.Scan(&expense_need_to_update.Price)

	case choice == 5:
		fmt.Print("Edit Description: " + expense_need_to_update.Description)
		descripton, _ := reader.ReadString('\n')
		expense_need_to_update.Description = descripton
	}

	/////////////////////////

	fmt.Println("You updated the Expense!")

	deleteExpense(id) // Delete the old copy of the updated expense

	return expense_need_to_update

}

func findExpenseByID(id int) Expense {

	var data Expense

	readExpense()

	for _, single_expense := range expenses_slice {
		if single_expense.ID == id {
			data = single_expense
		}
	}

	return data
}

func fileExist() bool {

	var true_or_false bool

	_, err := os.Stat(JSONFILENAME)

	if err != nil {

		if os.IsNotExist(err) {
			true_or_false = false
		} else {

			true_or_false = false
			fmt.Println("Error: ", err)
		}

	} else {
		true_or_false = true
	}

	return true_or_false

}
