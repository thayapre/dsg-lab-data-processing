package rawdataextractor

import (
	"strings"
)

var restaurantHeader []string = []string{
	"name",
	"node_id",
	"postcode",
	"city",
	"street",
	"housenumber",
	"cuisine",
	"outdoor_seating",
	"takeaway",
}

// Extracts the restaurants from a raw dataset and combines them with OSM data
func ExtractRestaurants(filepath string) ([][]string, error) {
	restaurantNodeIDs := make(map[string]string)

	records, err := ReadCSV(filepath)
	if err != nil {
		return nil, err
	}

	for i, record := range records {

		if i > 0 {
			nodeID := record[1][strings.LastIndex(record[1], "(")+1 : strings.LastIndex(record[1], ")")]
			restaurantNodeIDs[nodeID] = ""
		}
	}

	restaurants := [][]string{restaurantHeader}

	for nodeID, _ := range restaurantNodeIDs {
		restaurant, err := GetOSMRestaurantData(nodeID)
		if err != nil {
			return nil, err
		}

		temp_record := []string{
			restaurant.name, restaurant.nodeID,
			restaurant.postcode, restaurant.city, restaurant.street, restaurant.housenumber,
			restaurant.cuisine,
			restaurant.outdoorSeating,
			restaurant.takeaway,
		}

		restaurants = append(restaurants, temp_record)
	}

	return restaurants, nil
}
