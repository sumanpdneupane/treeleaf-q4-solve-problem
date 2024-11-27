package helpers

import (
	"q4-problem-solve/models"

	"github.com/google/uuid"
)

// ConvertCustomerProfiles updates the customer profile collection with new 128-bit UUIDs.
func ConvertCustomerProfiles(profiles map[int]models.CustomerProfile, oldCustomerID *int, newProfile models.CustomerProfile) map[int]models.CustomerProfile {
	// If oldCustomerID is provided, find and update the existing profile.
	if oldCustomerID != nil {
		if profile, exists := profiles[*oldCustomerID]; exists {
			// Create a new profile with the same data but a new UUID.
			newUUID := uuid.New().String()
			updatedProfile := models.CustomerProfile{
				OldCustomerID: profile.OldCustomerID,
				NewCustomerID: newUUID,
				Name:          profile.Name,
				Address:       profile.Address,
				RFIDNumber:    profile.RFIDNumber,
				JoinedDate:    profile.JoinedDate,
			}
			// Add the updated profile to the collection (old profiles remain immutable).
			profiles[profile.OldCustomerID] = updatedProfile
			return profiles
		}
	}

	// If oldCustomerID is nil, treat it as a new customer.
	newUUID := uuid.New().String()
	newProfile.NewCustomerID = newUUID
	newProfile.OldCustomerID = 0           // No old customer ID for new profiles.
	profiles[len(profiles)+1] = newProfile // Add the new profile to the collection.
	return profiles
}
