package interfaces

import (
	"hello_go/domain"
)

// OwnerManager interface for owner manager
type OwnerManager interface {
	GetOwners() map[string]*domain.Owner
	CreateOwner(ownerID string, name string, address *domain.Address) (*domain.Owner, error)
	GetOwner(ownerID string) *domain.Owner
	UpdateOwner(owner *domain.Owner) bool
	DeleteOwner(owner *domain.Owner) bool

	IsExisting(ownerID string) bool
}
