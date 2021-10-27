package rawdataextractor

var ratingHeader []string = []string{
	"rating_overall",
	"rating_service",
	"restaurant_price_range",
	"rating_location",
	"breakfast", "lunch", "regular_dinner", "upscale_dinner",
	"alone", "friends", "students", "work_colleagues", "partner",
	"takeaway",
	"student_discount",
	"wheelchair_accesible",
	"person_id",
	"rating_date",
	"restaurant_node_id",
}

// Extracts the ratings from a raw dataset
func ExtractRatings(filepath string) ([][]string, error) {
	records, err := ReadCSV(filepath)
	if err != nil {
		return nil, err
	}

	for i, record := range records {
		if i > 0 {
			temp_record := record[2:18]
			temp_record = append(temp_record, record[len(record)-4], record[len(record)-2])

			records[i] = temp_record
		}
	}

	records[0] = ratingHeader

	return records, nil
}
