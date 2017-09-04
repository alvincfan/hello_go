package interfaces

import "hello_go/domain"

// DealershipManager interface for dealership manager
type DealershipManager interface {
	GetAllDealership() map[string]*domain.Dealership
	AddDealership(dealership *domain.Dealership) error
	GetDealership(dealershipID string) *domain.Dealership
	RemoveDealership(dealershipID string) error

	AddGroundInventory(dealershipID string, ground *domain.GroundTransportation) error
	GetGroundInventory(dealershipID string, serialNumber string) (*domain.GroundTransportation, error)
	MoveGroundInventory(dealershipID string, serialNumber string) error
	SoldGroundInventory(dealershipID string, serialNumber string, owner *domain.Owner) error
}
