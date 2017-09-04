package domain

// Address holds address
type Address struct {
	Geolocation
	Zipcode        string
	MailingAddress string
}
