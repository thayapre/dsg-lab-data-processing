package main

import (
	"log"
	"os"

	"github.com/DS21t/dsg-lab-data-processing/internal/rawdataextractor"
)

func main() {
	// Extract persons from entrypoint
	persons, err := rawdataextractor.ExtractPersons(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Extract ratings from entrypoint and plain ratings
	ratings, err := rawdataextractor.ExtractRatings(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	plainRatings, err := rawdataextractor.ExtractRatings(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Combine ratings from both sources and remove header from the second source
	ratings = append(ratings, plainRatings[1:]...)

	err = rawdataextractor.WriteCSV(persons, os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	err = rawdataextractor.WriteCSV(ratings, os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
}
