package main

import (
	"fmt"
	"q4-problem-solve/helpers"
	"q4-problem-solve/models"
)

func main() {
	// Initialize a collection of customer profiles using a map.
	customerProfiles := map[int]models.CustomerProfile{
		1001: {OldCustomerID: 1001, NewCustomerID: "", Name: "Alice", Address: "123 Main St", RFIDNumber: 2001, JoinedDate: "2024-01-01"},
		1002: {OldCustomerID: 1002, NewCustomerID: "", Name: "Bob", Address: "456 Elm St", RFIDNumber: 2002, JoinedDate: "2024-01-02"},
	}

	// Convert an existing customer profile (update with UUID).
	oldID := 1001
	customerProfiles = helpers.ConvertCustomerProfiles(customerProfiles, &oldID, models.CustomerProfile{})

	// Add a new customer profile.
	newProfile := models.CustomerProfile{
		Name:       "Charlie",
		Address:    "789 Oak St",
		RFIDNumber: 2003,
		JoinedDate: "2024-11-27",
	}
	customerProfiles = helpers.ConvertCustomerProfiles(customerProfiles, nil, newProfile)

	// Print the updated customer profiles.
	for id, profile := range customerProfiles {
		fmt.Printf("Profile ID: %d, Old ID: %d, New UUID: %s, Name: %s, Address: %s, RFID: %d, Joined Date: %s\n",
			id, profile.OldCustomerID, profile.NewCustomerID, profile.Name, profile.Address, profile.RFIDNumber, profile.JoinedDate)
	}
}
