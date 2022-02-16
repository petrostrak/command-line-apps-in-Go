package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statistical function
type statsFunc func([]float64) float64

func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// csvToFloat receives an io.Reader interface representing the source of CSV data
// and an int representing the column to extract data from.
//
// The program assumes the user will input the column starting from (1)
// ass it's more natural for users to understand.
func csvToFloat(r io.Reader, column int) ([]float64, error) {

	// create the csv Reader used to read in data from CSV files
	cr := csv.NewReader(r)

	// adjusting for 0 based index
	column--

	// read in all csv data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	var data []float64

	// looping through all records
	for i, row := range allData {
		if i == 0 {
			continue
		}

		// checking number of columns in CSV file
		if len(row) <= column {
			// file does not have that many columns
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}

		// try to convert data read into float number
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	// return the slice of float64 and nil error
	return data, nil
}
