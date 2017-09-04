package domain

// Dealership holds dealership object
type Dealership struct {
	Address
	Name            string
	DealershipID    string
	GroundInventory map[string]*GroundTransportation
}
