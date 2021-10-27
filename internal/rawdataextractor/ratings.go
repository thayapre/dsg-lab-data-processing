package rawdataextractor

import (
	"strings"
)

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
	"restaurant_name",
	"restaurant_node_id",
	"restaurant_postcode",
	"restaurant_city",
	"restaurant_street",
	"restaurant_housenumber",
	"restaurant_cuisine",
	"restaurant_outdoor_seating",
	"restaurant_takeaway",
}

// Extracts the ratings from a raw dataset and combines them with OSM data
func ExtractRatings(filepath string) ([][]string, error) {
	records, err := ReadCSV(filepath)
	if err != nil {
		return nil, err
	}

	for i, record := range records {
		if i > 0 {
			temp_record := record[2:18]
			temp_record = append(temp_record, record[len(record)-4], record[len(record)-2])

			nodeID := record[1][strings.LastIndex(record[1], "(")+1 : strings.LastIndex(record[1], ")")]

			restaurant, err := GetOSMRestaurantData(nodeID)
			if err != nil {
				return nil, err
			}

			temp_record = append(temp_record,
				restaurant.name, restaurant.nodeID,
				restaurant.postcode, restaurant.city, restaurant.street, restaurant.housenumber,
				restaurant.cuisine,
				restaurant.outdoorSeating,
				restaurant.takeaway,
			)

			records[i] = temp_record
		}
	}

	records[0] = ratingHeader

	return records, nil
}
