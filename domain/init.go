package domain

// NewGroundTransportation creates a new ground transportation
func NewGroundTransportation(manufacture *Manufacture, location Geolocation, serialNumber string) *GroundTransportation {
	ground := &GroundTransportation{
		Location: location,
		SerialNumber: serialNumber,
		Manufacture: manufacture,
	}
	return ground
}

// NewFactory creates a new factory
func NewFactory(name string, yearInService int32, address *Address) *Manufacture {
	factory := &Manufacture{
		Name:          name,
		YearInService: yearInService,
		Address:       *address,
	}
	return factory
}

// NewDealership creates a new dealership
func NewDealership(dealershipID string, name string, address *Address) *Dealership {
	dealership := &Dealership{
		DealershipID:    dealershipID,
		Name:            name,
		Address:         *address,
		GroundInventory: make(map[string]*GroundTransportation),
	}
	return dealership
}

// NewOwner creates a new owner
func NewOwner(ownerID string, name string, address *Address) *Owner {
	owner := &Owner{
		Name:           name,
		Address:        *address,
		OwnerID:        ownerID,
		Transportation: make([]string, 1),
	}
	return owner
}
