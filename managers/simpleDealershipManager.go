package managers

import (
	"fmt"
	"hello_go/domain"
	"hello_go/interfaces"
)

// SimpleDealershipManager holds an implementation of dealership manager
type SimpleDealershipManager struct {
	dealerships map[string]*domain.Dealership
}

// NewSimpleDealershipManager creates a dealership manager
func NewSimpleDealershipManager() interfaces.DealershipManager {
	dealershipManager := &SimpleDealershipManager{
		dealerships: make(map[string]*domain.Dealership),
	}
	return dealershipManager
}

//GetAllDealership return dealership
func (sD *SimpleDealershipManager) GetAllDealership() map[string]*domain.Dealership {
	return sD.dealerships
}

// GetDealership get dealership
func (sD *SimpleDealershipManager) GetDealership(dealershipID string) *domain.Dealership {
	if dealership, ok := sD.dealerships[dealershipID]; ok {
		return dealership
	}
	return nil
}

// AddDealership add dealership
func (sD *SimpleDealershipManager) AddDealership(dealership *domain.Dealership) error {
	dealershipID := sD.GetDealership(dealership.DealershipID)
	if dealershipID != nil {
		return fmt.Errorf("dealerships %v already added", dealershipID)
	}

	sD.dealerships[dealership.DealershipID] = dealership
	return nil
}

// RemoveDealership remove dealership
func (sD *SimpleDealershipManager) RemoveDealership(dealershipID string) error {
	dealership := sD.GetDealership(dealershipID)
	if dealership == nil {
		return fmt.Errorf("dealerships not found %v", dealershipID)
	}
	delete(sD.dealerships, dealershipID)
	return nil
}

// AddGroundInventory add ground inventory
func (sD *SimpleDealershipManager) AddGroundInventory(dealershipID string, ground *domain.GroundTransportation) error {
	dealership := sD.GetDealership(dealershipID)
	if dealership == nil {
		return fmt.Errorf("AddGroundInventory failed, dealerships not found based on dealershipID %v for Add", dealershipID)
	}

	if found, ok := dealership.GroundInventory[ground.SerialNumber]; ok {
		return fmt.Errorf("Serial Number %v already in dealership ground inventory", found.SerialNumber)
	}
	dealership.GroundInventory[ground.SerialNumber] = ground

	return nil
}

// GetGroundInventory get a ground transportation
func (sD *SimpleDealershipManager) GetGroundInventory(dealershipID string, serialNumber string) (*domain.GroundTransportation, error) {

	dealership := sD.GetDealership(dealershipID)

	if dealership == nil {
		return nil, fmt.Errorf("dealerships not found based on dealershipID %v", dealershipID)
	}

	if found, ok := dealership.GroundInventory[serialNumber]; ok {
		return found, nil

	}
	return nil, fmt.Errorf("Ground Inventory not found based on serial number %v", serialNumber)
}

// MoveGroundInventory move a ground inventory to another dealership
func (sD *SimpleDealershipManager) MoveGroundInventory(dealershipID string, serialNumber string) error {

	if _, error := sD.GetGroundInventory(dealershipID, serialNumber); error != nil {
		return error
	}
	return nil
}

// SoldGroundInventory solde a ground inventory to an owner
func (sD *SimpleDealershipManager) SoldGroundInventory(dealershipID string, serialNumber string, owner *domain.Owner) error {
	found, error := sD.GetGroundInventory(dealershipID, serialNumber)
	if error != nil {
		return error
	}
	fmt.Printf("Update ground inventory serial number %v owner to %v\n", found.SerialNumber, owner.Name)
	found.Owner = owner
	fmt.Printf("Remove serial number %v from dealership %v inventory\n", serialNumber, dealershipID)
	delete(sD.GetDealership(dealershipID).GroundInventory, serialNumber)
	return nil
}
