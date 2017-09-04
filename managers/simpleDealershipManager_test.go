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

var _ = Describe("Dealershipmanager", func() {
	var dm interfaces.DealershipManager
	var om interfaces.OwnerManager
	var owner *domain.Owner
	var dealership *domain.Dealership
	var ground *domain.GroundTransportation
	var manufacture *domain.Manufacture

	BeforeEach(func() {
		dm = managers.NewSimpleDealershipManager()
		om = mock_interfaces.NewMockOwnerManager(mockCtrl)
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

	Context("when calling get dealership", func() {
		Context("when dealership ID is not found", func() {
			It("should returns nil", func() {
				Expect(dm.GetDealership(dealership.DealershipID)).To(BeNil())

				By("add dealership to management")
				Expect(dm.AddDealership(dealership)).To(BeNil())

				By("add duplicate dealership to management")
				Expect(dm.AddDealership(dealership)).ToNot(BeNil())

				By(" remove dealership from management")
				Expect(dm.RemoveDealership(dealership.DealershipID)).To(BeNil())

				By("remove same dealership from management")
				Expect(dm.RemoveDealership(dealership.DealershipID)).ToNot(BeNil())
			})
		})
	})

	Context("when calling add ground inventory", func() {
		Context(" when dealership ID is not found", func() {
			It("should return error", func() {
				Expect(dm.AddGroundInventory(dealership.DealershipID, ground)).ToNot(BeNil())

				By("add dealership to management")
				Expect(dm.AddDealership(dealership)).To(BeNil())

				By("add transportation to dealership")
				Expect(dm.AddGroundInventory(dealership.DealershipID, ground)).To(BeNil())

				By("add duplcate transportation to dealership")
				Expect(dm.AddGroundInventory(dealership.DealershipID, ground)).ToNot(BeNil())

				By(" sold transportation to owner")
				Expect(dm.SoldGroundInventory(dealership.DealershipID, ground.SerialNumber, owner)).To(BeNil())
				Expect(ground.Owner.OwnerID).To(Equal(owner.OwnerID))

				_, err := dm.GetGroundInventory(dealership.DealershipID, ground.SerialNumber)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	PContext("when calling move ground inventory", func() {
		It("should move inventory to dealership", func() {

		})
	})
})
