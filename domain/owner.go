package domain

// Owner holds owner object
type Owner struct {
	Address
	OwnerID        string
	Name           string
	Transportation []string
}
