package rawdataextractor

import (
	overpass "github.com/serjvanilla/go-overpass"
)

type Restaurant struct {
	name           string
	postcode       string
	city           string
	street         string
	housenumber    string
	cuisine        string
	outdoorSeating string
	takeaway       string
	nodeID         string
}

// Queries the restaurant specific OSM data for the given node ID.
func GetOSMRestaurantData(nodeID string) (Restaurant, error) {
	restaurant := Restaurant{nodeID: nodeID}

	result, err := overpass.DefaultClient.Query("[out:json];node(" + nodeID + ");out body;")
	if err != nil {
		return Restaurant{}, err
	}

	for _, node := range result.Nodes {
		for key, value := range node.Tags {
			switch key {
			case "name":
				restaurant.name = value
			case "addr:postcode":
				restaurant.postcode = value
			case "addr:city":
				restaurant.city = value
			case "addr:street":
				restaurant.postcode = value
			case "addr:housenumber":
				restaurant.housenumber = value
			case "cuisine":
				restaurant.cuisine = value
			case "outdoor_seating":
				restaurant.outdoorSeating = value
			case "takeaway":
				restaurant.takeaway = value
			}
		}
	}

	return restaurant, nil
}
