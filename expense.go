package main

import (
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

func (exp Expense) addExpense() string {

	filename := "expense.json"

	var expenses []Expense

	expense := Expense{ID: exp.ID, Date: exp.Date, Category: exp.Category, Item: exp.Item, Price: exp.Price, Description: exp.Description}

	expenses = append(expenses, expense)

	if !fileExist(filename) {
		writeExpense(expenses, filename)
	} else {
		appendExpense(filename, expense)
	}

	return "You successfully added an expense!\n"

}

func readExpense(filename string) ([]Expense, error) {

	file, err := os.Open(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}

		return nil, err
	}

	defer file.Close()

	var expenses []Expense

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&expenses)

	if err != nil {
		return nil, err
	}

	return expenses, nil

}

func writeExpense(expenses []Expense, filename string) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(expenses)

	if err != nil {
		return err
	}

	return nil
}

func appendExpense(filename string, newExpense Expense) error {

	expenses, err := readExpense(filename)

	if err != nil {
		return err
	}

	expenses = append(expenses, newExpense)

	err = writeExpense(expenses, filename)
	if err != nil {
		return err
	}

	return nil
}

func expensesReport() {
	exp, _ := readExpense("expense.json")

	var data [][]string

	for _, d := range exp {

		dat := []string{fmt.Sprintf("%d", d.ID), d.Date, d.Category, d.Item, fmt.Sprintf("%.2f", d.Price), d.Description}

		data = append(data, dat)
	}

	x := []string{"ID", "Date", "Category", "Item", "Price", "Description"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(x)

	for _, row := range data[0:] {

		table.Append(row)
	}

	table.Render()

}

func fileExist(filename string) bool {

	var v bool

	_, err := os.Stat(filename)

	if err != nil {

		if os.IsNotExist(err) {
			v = false
		} else {

			v = false
			fmt.Println("Error: ", err)
		}

	} else {
		v = true
	}

	return v

}
