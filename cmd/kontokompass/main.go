package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/zachpanter/kontokompass/internal/storage"
	_ "github.com/zachpanter/kontokompass/internal/storage"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {

	ctx := context.Background()
	folderPath := "/Users/blackjack/Desktop/bank_data" // Replace with the actual path to your folder

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	allData := [][]string{} // A slice to store data from all CSV files

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".csv" {
			continue // Skip directories and non-CSV files
		}

		filePath := filepath.Join(folderPath, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			log.Printf("Error opening %s: %v\n", filePath, err)
			continue // Skip to the next file if there's an error
		}

		r := csv.NewReader(f)
		rows, err := r.ReadAll()
		if err != nil {
			log.Printf("Error reading %s: %v\n", filePath, err)
			f.Close()
			continue // Skip to the next file if there's an error
		}
		rows = rows[1:]
		allData = append(allData, rows...)
		closeErr := f.Close()
		if closeErr != nil {
			return
		}
	}

	// Now you have all the CSV data in memory
	fmt.Println(allData)

	importToDB(ctx, allData)

}

func importToDB(ctx context.Context, allData [][]string) {
	queries := storage.OpenDBPool(ctx)

	for index, row := range allData {
		var params storage.InsertTransactionParams

		// Acct# is index 0

		// Date
		postDateString := row[1]
		layout := "01/01/2006" // Layout must match the format of your date string

		dateVal, dateParseErr := time.Parse(layout, postDateString)
		if dateParseErr != nil {
			fmt.Println("Error parsing date:", dateParseErr)
			return
		}
		params.Postdate = dateVal

		// Check is index 2

		// Description
		params.Description = row[3]

		// Debit
		debit, debitParseErr := strconv.ParseFloat(row[4], 64)
		if debitParseErr != nil {
			params.Debit.Float64 = 0.0
			params.Debit.Valid = false
		} else {
			params.Debit.Float64 = debit
			params.Debit.Valid = true
		}

		// Credit
		credit, creditParseErr := strconv.ParseFloat(row[5], 64)
		if creditParseErr != nil {
			params.Credit.Float64 = 0.0
			params.Credit.Valid = false
		} else {
			params.Credit.Float64 = credit
			params.Credit.Valid = true
		}

		// Status is index 6

		// Balance
		balance, balanceParseErr := strconv.ParseFloat(row[7], 32)
		if balanceParseErr != nil {
			log.Panic(balanceParseErr)
		} else {
			params.Balance = float32(balance)
		}

		params.ClassificationText = allData[index][8]
		insertTransactionErr := queries.InsertTransaction(ctx, params)
		if insertTransactionErr != nil {
			return
		}
	}
}
