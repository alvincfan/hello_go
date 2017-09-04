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
	found := oM.GetOwner(owner.OwnerID)
	if found == nil {
		return false
	}
	oM.owners[owner.OwnerID] = owner
	return true
}

// DeleteOwner delete existing owner
func (oM *SimpleOwnerManager) DeleteOwner(owner *domain.Owner) bool {
	found := oM.GetOwner(owner.OwnerID)
	if found == nil {
		return false
	}
	delete(oM.owners, found.OwnerID)
	return true
}
