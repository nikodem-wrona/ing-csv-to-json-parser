package transaction

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/gocarina/gocsv"
)

func ProcessCSV(filename string) []*RawTransaction {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = ';'
		return r
	})

	transactionsFile, err := os.OpenFile(filename, os.O_RDWR, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer func() {
		transactionsFile.Close()
	}()

	transactions := []*RawTransaction{}

	err = gocsv.UnmarshalFile(transactionsFile, &transactions)

	if err != nil {
		panic(err)
	}

	return transactions
}

func CreateJSONOutput(rawTransactions []*RawTransaction) {
	transactions := []Transaction{}

	for _, t := range rawTransactions {
		transactions = append(transactions, t.ProcessTransaction())
	}

	file, err := json.MarshalIndent(transactions, "", " ")

	if err != nil {
		panic("Failed to encode to json")
	}

	err = ioutil.WriteFile("output.json", file, 0644)

	if err != nil {
		panic("Failed to save transactions to file")
	}
}
