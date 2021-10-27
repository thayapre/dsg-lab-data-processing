package rawdataextractor

var personHeader []string = []string{
	"age",
	"gender",
	"kids",
	"alcohol",
	"family_status",
	"occupation",
	"diet",
	"intolerance_gluten", "intolerance_sorbitol", "intolerance_lactose", "intolerance_fructose", "intolerance_histamine", "intolerance_sucrose", "intolerance_none",
	"preferred_drink",
	"sports_state",
	"cooking_frequency",
	"person_id",
}

// Extracts the persons from a raw dataset
func ExtractPersons(filepath string) ([][]string, error) {
	records, err := ReadCSV(filepath)
	if err != nil {
		return nil, err
	}

	for index, record := range records {
		temp_record := record[18:35]
		temp_record = append(temp_record, record[len(record)-4], record[len(record)-2])

		records[index] = temp_record
	}

	records[0] = personHeader

	return records, nil
}
