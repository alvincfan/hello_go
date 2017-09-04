package managers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/twinj/uuid"
	"hello_go/domain"
	"hello_go/interfaces"
	"hello_go/interfaces/mock_interfaces"
	"hello_go/managers"
)

var _ = Describe("Ownermanager", func() {
	var dm interfaces.DealershipManager
	var om interfaces.OwnerManager
	var owner *domain.Owner
	var dealership *domain.Dealership
	var ground *domain.GroundTransportation
	var manufacture *domain.Manufacture

	BeforeEach(func() {
		dm = mock_interfaces.NewMockDealershipManager(mockCtrl)
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

	Context("when calling get owner", func() {
		Context("when not in owner system", func() {
			It("should return nil", func() {
				Expect(om.GetOwner(owner.OwnerID)).To(BeNil())

				By("add owner to system")
				_, err := om.CreateOwner(owner.OwnerID, owner.Name, &owner.Address)
				Expect(err).To(BeNil())

				By("add duplicate owner to system")
				_, err = om.CreateOwner(owner.OwnerID, owner.Name, &owner.Address)
				Expect(err).ToNot(BeNil())

				owner.MailingAddress = "update mailing address"
				Expect(om.UpdateOwner(owner)).To(BeTrue())
				Expect(owner.MailingAddress).To(Equal(om.GetOwner(owner.OwnerID).MailingAddress))

				By(" delete owner from system", func() {
					Expect(om.DeleteOwner(owner)).To(BeTrue())
					By("delete owner not in system")
					Expect(om.DeleteOwner(owner)).To(BeFalse())
				})
			})
		})
	})
})
