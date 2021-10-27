package rawdataextractor

import (
	"encoding/csv"
	"os"
)

// Reads the records from a CSV file
func ReadCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// Writes the records to a CSV file
func WriteCSV(records [][]string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	w.WriteAll(records)

	return nil
}
