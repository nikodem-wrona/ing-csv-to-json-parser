package main

import (
	"fmt"

	"github.com/nikodemwrona/ing_parser/pkg/input"
	"github.com/nikodemwrona/ing_parser/pkg/transaction"
)

func main() {
	arguments := input.ProcessArguments()
	transactions := transaction.ProcessCSV(arguments.Filename)

	transaction.CreateJSONOutput(transactions)

	fmt.Println("Done.")
}
