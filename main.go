package main

import (
	"fmt"
	"github.com/twinj/uuid"
	"hello_go/domain"
	"hello_go/interfaces"
	"hello_go/managers"
	"hello_go/usecases"
)

func initialized() (interfaces.OwnerManager, interfaces.DealershipManager) {
	om := managers.NewSimpleOwnerManager()
	dm := managers.NewSimpleDealershipManager()
	return om, dm
}

func main() {
	defaultAddress := domain.Address{
		MailingAddress: "default",
		Geolocation: domain.Geolocation{
			Longitude: 0,
			Latitude:  0,
		},
	}
	ownerManager, dealershipManager := initialized()

	owner := domain.NewOwner(uuid.NewV4().String(), "me", &defaultAddress)
	f := domain.NewFactory("myhome", 200, &defaultAddress)
	g := domain.NewGroundTransportation(f, domain.Geolocation{
		Latitude:  100,
		Longitude: 300}, "abcdef")
	d := domain.NewDealership(uuid.NewV4().String(), "mydealer", &defaultAddress)

	err := dealershipManager.AddDealership(d)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = usecases.AddGroundInventory(dealershipManager, d.DealershipID, g)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prevOwner := usecases.CheckGroundTransportationOwner(ownerManager, g.SerialNumber)
	if prevOwner != nil {
		fmt.Println("Transporation Serial Number %v has previous owner %v", g.SerialNumber, prevOwner.Name)
	} else {
		fmt.Println("Transporation Serial Number %v has no previous owner recorded", g.SerialNumber)
	}

	err = usecases.SoldGroundTransportation(dealershipManager, ownerManager, d, g, owner)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	owner.MailingAddress = "mymailingaddress"
	d.MailingAddress = "1 infinity road"
	f.MailingAddress = "china"

	fmt.Printf("Owner %v address %v transporataion serial number %v\n",
		ownerManager.GetOwner(owner.OwnerID).Name,
		ownerManager.GetOwner(owner.OwnerID).MailingAddress,
		ownerManager.GetOwner(owner.OwnerID).Transportation,
	)

	fmt.Printf("Dealer %v address %v\n",
		dealershipManager.GetDealership(d.DealershipID).Name,
		dealershipManager.GetDealership(d.DealershipID).MailingAddress)

	fmt.Printf("Ground Transportation %v Owner %v\n",
		g.SerialNumber, g.Owner.Name)
	fmt.Printf("default address at main %v\n", defaultAddress)

	err = usecases.BuyGroundTransportationFromOwner(dealershipManager, ownerManager, d, g, owner)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	prevOwner = usecases.CheckGroundTransportationOwner(ownerManager, g.SerialNumber)
	if prevOwner != nil {
		fmt.Println("Transporation Serial Number %v has previous owner %v", g.SerialNumber, prevOwner.Name)
	} else {
		fmt.Println("Transporation Serial Number %v has no previous owner recorded", g.SerialNumber)
	}

	usecases.SoldGroundTransportation(dealershipManager, ownerManager, d, g, owner)

	fmt.Printf("Owner %v address %v transporataion serial number %v\n",
		ownerManager.GetOwner(owner.OwnerID).Name,
		ownerManager.GetOwner(owner.OwnerID).MailingAddress,
		ownerManager.GetOwner(owner.OwnerID).Transportation,
	)

}
