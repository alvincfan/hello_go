package managers

import (
	"fmt"
	"hello_go/domain"
	"hello_go/interfaces"
)

// SimpleOwnerManager holds an implementation of owner manager
type SimpleOwnerManager struct {
	owners map[string]*domain.Owner
}

// NewSimpleOwnerManager creates an owner manager
func NewSimpleOwnerManager() interfaces.OwnerManager {
	ownerManager := &SimpleOwnerManager{
		owners: make(map[string]*domain.Owner),
	}
	return ownerManager
}

// GetOwners return all owners
func (oM *SimpleOwnerManager) GetOwners() map[string]*domain.Owner {
	return oM.owners
}

// CreateOwner create owner
func (oM *SimpleOwnerManager) CreateOwner(ownerID string, name string, address *domain.Address) (*domain.Owner, error) {

	if oM.GetOwner(ownerID) != nil {
		return nil, fmt.Errorf("OwnerID %v already existed", ownerID)
	}

	owner := &domain.Owner{
		OwnerID: ownerID,
		Name:    name,
		Address: *address,
	}

	oM.owners[ownerID] = owner
	return owner, nil
}

// GetOwner get owner
func (oM *SimpleOwnerManager) GetOwner(owerID string) *domain.Owner {
	if owner, ok := oM.owners[owerID]; ok {
		return owner
	}
	return nil
}

// UpdateOwner update existing owner information
func (oM *SimpleOwnerManager) UpdateOwner(owner *domain.Owner) bool {
	if oM.IsExisting(owner.OwnerID) {
		oM.owners[owner.OwnerID] = owner
		return true
	}
	return false
}

// DeleteOwner delete existing owner
func (oM *SimpleOwnerManager) DeleteOwner(owner *domain.Owner) bool {
	if oM.IsExisting(owner.OwnerID) {
		delete(oM.owners, owner.OwnerID)
		return true
	}
	return false

}

// IsExisting owner is in the system
func (oM *SimpleOwnerManager) IsExisting(ownerID string) bool {
	found := oM.GetOwner(ownerID)
	if found == nil {
		return false
	}
	return true
}