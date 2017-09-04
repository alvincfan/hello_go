package domain

// GroundTransportation holds ground transportation object
type GroundTransportation struct {
	Manufacture  *Manufacture
	Owner        *Owner
	Location     Geolocation
	SerialNumber string
}
