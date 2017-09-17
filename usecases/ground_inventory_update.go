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

// SoldGroundTransportation update dealership manager and owner manager when sold a ground transporatation to an owner
func SoldGroundTransportation(dealershipMgr interfaces.DealershipManager, ownerMgr interfaces.OwnerManager, dealership *domain.Dealership, ground *domain.GroundTransportation, owner *domain.Owner) error {

	owner.Transportation = append(owner.Transportation, ground.SerialNumber)

	if !UpdateOwnerInOwnerManagement(ownerMgr, owner) {
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

// UpdateOwnerInOwnerManagement finds out if owner is in the system
func UpdateOwnerInOwnerManagement(ownerMgr interfaces.OwnerManager, owner *domain.Owner) bool {
	if ownerMgr.IsExisting(owner.OwnerID) {
		// repeat customer!  we are doing something right for owner to comes back for more!
		fmt.Println("Repeat customer!  we are doing something right!")
	} else {
		// owner not in owner manager system, add for future advertisement! ahaha
		fmt.Println("add customer to owner system for advertisement!")
		ownerMgr.CreateOwner(owner.OwnerID, owner.Name, &owner.Address)
	}

	return ownerMgr.UpdateOwner(owner)
}

// BuyGroundTransportationFromOwner buys back ground transporatation
func BuyGroundTransportationFromOwner(dealershipMgr interfaces.DealershipManager, ownerMgr interfaces.OwnerManager, dealership *domain.Dealership, ground *domain.GroundTransportation, owner *domain.Owner) error {
	ground.Owner = nil
	newTransportation := []string{}

	// Update owner transportation
	for _, t := range owner.Transportation {
		if t == ground.SerialNumber {
			continue
		}
		newTransportation = append(newTransportation, t)
	}
	owner.Transportation = newTransportation

	if !UpdateOwnerInOwnerManagement(ownerMgr, owner) {
		return fmt.Errorf("Failed to update owner %v", owner)
	}

	return dealershipMgr.AddGroundInventory(dealership.DealershipID, ground)
}