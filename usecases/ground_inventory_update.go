package usecases

import (
	"fmt"
	"hello_go/domain"
	"hello_go/interfaces"
)

// AddGroundInventory adds ground inventory to a dealership
func AddGroundInventory(dealershipMgr interfaces.DealershipManager, dealershipID string, ground *domain.GroundTransportation) error {
	found, ID := CheckGroundTransportationInDealership(dealershipMgr, ground.SerialNumber)
	if found != nil {
		return fmt.Errorf("Ground Serial Number %v already under dealerships %v", found.SerialNumber, ID)
	}

	return dealershipMgr.AddGroundInventory(dealershipID, ground)
}

// AddNewDealership adds a dealership to dealership manager
func AddNewDealership(dealershipMgr interfaces.DealershipManager, dealership *domain.Dealership) error {
	return dealershipMgr.AddDealership(dealership)
}

// SoldGroundTransportation update dealership manager and owner manager when sold a ground transporatation to an owner
func SoldGroundTransportation(dealershipMgr interfaces.DealershipManager, ownerMgr interfaces.OwnerManager, dealership *domain.Dealership, ground *domain.GroundTransportation, owner *domain.Owner) error {

	owner.Transportation = append(owner.Transportation, ground.SerialNumber)

	existingOwner := ownerMgr.GetOwner(owner.OwnerID)
	if existingOwner == nil {
		// owner not in owner manager system, add for future advertisement! ahaha
		fmt.Println("add customer to owner system for advertisement!")
		ownerMgr.CreateOwner(owner.OwnerID, owner.Name, &owner.Address)
	} else {
		// repeat customer!  we are doing something right for owner to comes back for more!
		fmt.Println("Repeat customer!  we are doing something right!")
	}

	if !ownerMgr.UpdateOwner(owner) {
		return fmt.Errorf("Failed to update owner %v", owner)
	}

	error := dealershipMgr.SoldGroundInventory(dealership.DealershipID, ground.SerialNumber, owner)
	if error != nil {
		return error
	}

	return nil
}

// CheckGroundTransportationInDealership finds out a ground transportation in a dealership
func CheckGroundTransportationInDealership(dealershipMgr interfaces.DealershipManager, serialNumber string) (*domain.GroundTransportation, string) {
	for ID, dealership := range dealershipMgr.GetAllDealership() {
		if found, ok := dealership.GroundInventory[serialNumber]; ok {
			return found, ID
		}
	}
	return nil, ""
}

// CheckGroundTransportationOwner finds out owner of a ground transportation
func CheckGroundTransportationOwner(ownerMgr interfaces.OwnerManager, serialNumber string) *domain.Owner {
	for _, owner := range ownerMgr.GetOwners() {
		for _, num := range owner.Transportation {
			if num == serialNumber {
				return owner
			}
		}
	}
	return nil
}
