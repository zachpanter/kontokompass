package loader

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

func LoadCSV(ctx context.Context) {
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
}
