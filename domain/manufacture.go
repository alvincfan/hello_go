package domain

// Manufacture holds manufacture object
type Manufacture struct {
	Address
	ManufactureID string
	Name          string
	YearInService int32
}
