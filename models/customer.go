package models

// CustomerProfile represents the structure of a customer profile.
type CustomerProfile struct {
	OldCustomerID int    // Old 4-byte integer customer ID (can be 0 for new customers)
	NewCustomerID string // New 128-bit UUID customer ID
	Name          string // Customer name
	Address       string // Customer address
	RFIDNumber    int64  // RFID number (8-byte integer)
	JoinedDate    string // Customer joined date
}
