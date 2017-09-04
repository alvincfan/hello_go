package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/twinj/uuid"
	"hello_go/domain"
	"hello_go/interfaces"
	"hello_go/managers"
	"hello_go/usecases"
)

var _ = Describe("GroundInventoryUpdate", func() {
	var dm interfaces.DealershipManager
	var om interfaces.OwnerManager
	var owner *domain.Owner
	var dealership *domain.Dealership
	var ground *domain.GroundTransportation
	var manufacture *domain.Manufacture

	BeforeEach(func() {
		dm = managers.NewSimpleDealershipManager()
		om = managers.NewSimpleOwnerManager()
	})

	JustBeforeEach(func() {
		address := &domain.Address{
			MailingAddress: "test address",
			Zipcode:        "00000",
		}
		owner = domain.NewOwner(
			uuid.NewV4().String(),
			"testuser",
			address,
		)
		dealership = domain.NewDealership(
			uuid.NewV4().String(),
			"test dealership",
			address,
		)
		manufacture = domain.NewFactory(
			"test factory",
			2017,
			address,
		)
		location := domain.Geolocation{
			Latitude:  0,
			Longitude: 0,
		}
		ground = domain.NewGroundTransportation(
			manufacture, location, "test serial number",
		)

	})

	Context("when calling sold ground inventory to a dealership", func() {
		PIt("should update owner and dealership system", func() {
			Expect(usecases.SoldGroundTransportation(dm, om, dealership, ground, owner)).To(BeNil())

		})
	})

})
