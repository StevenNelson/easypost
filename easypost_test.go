package easypost

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSetup(t *testing.T) {
	EasyPostApi["Key"] = "x4NQEBqeFapSgb48kyWVSA"
}

func TestNewAddress(t *testing.T) {
	testFromAddress, err := NewAddress(&Address{
		Name:    "Steven Nelson",
		Street1: "1111 Main St",
		Street2: "Suite 204",
		City:    "Conway",
		State:   "AR",
		Zip:     "72032",
		Country: "US",
		Phone:   "5017335509",
		Email:   "stevenrnelson@gmail.com",
	})
	if !HasPrefix(testFromAddress.Id, "adr_") || err != nil {
		t.Log(testFromAddress.Id)
		t.Log(testFromAddress.Name)
		t.Fatal("creating test address failed")
	}
}

func TestRetrieveAddress(t *testing.T) {
	sentAddress := Address{
		Name:    "Steven Nelson",
		Street1: "1111 Main St",
		Street2: "Suite 204",
		City:    "Conway",
		State:   "AR",
		Zip:     "72032",
		Country: "US",
		Phone:   "5017335509",
		Email:   "stevenrnelson@gmail.com",
	}
	fromAddress, err := NewAddress(&sentAddress)
	if err != nil {
		t.Fatal("address couldn't be retrieved because it couldn't be created")
	}
	newAddress, err := RetrieveAddress(fromAddress.Id)
	if err != nil {
		t.Log(fromAddress.Id)
		t.Fatal(err)
	}
	if newAddress.Name != fromAddress.Name &&
		newAddress.Street1 != fromAddress.Street1 {
		newaddr, _ := json.Marshal(newAddress)
		sentaddr, _ := json.Marshal(fromAddress)
		t.Log(string(newaddr[:len(newaddr)]))
		t.Log(string(sentaddr[:len(sentaddr)]))
		t.Fatal("addresses don't match")
	}

}

func TestNewParcel(t *testing.T) {
	parcel, err := NewParcel(&Parcel{
		Length: 17.5,
		Width:  12,
		Height: 15.5,
		Weight: 124,
	})
	if !HasPrefix(parcel.Id, "prcl_") || err != nil {
		t.Fatal("creating test parcel failed")
	}
}

func TestRetrieveParcel(t *testing.T) {
	parcel, err := NewParcel(&Parcel{
		Length: 17.5,
		Width:  12,
		Height: 15.5,
		Weight: 124,
	})
	if !HasPrefix(parcel.Id, "prcl_") || err != nil {
		t.Fatal("parcel couldn't be retrieved because it couldn't be created")
	}

	newParcel, err := RetrieveParcel(parcel.Id)
	if err != nil {
		t.Log(parcel.Id)
		t.Fatal("couldn't retrieve parcel")
	}
	if parcel.Length != newParcel.Length || parcel.Width != newParcel.Width ||
		parcel.Height != newParcel.Height || parcel.Weight != newParcel.Weight {
		jsonParcel, _ := json.Marshal(parcel)
		jsonNewParcel, _ := json.Marshal(newParcel)
		t.Log(string(jsonParcel[:len(jsonParcel)]))
		t.Log(string(jsonNewParcel[:len(jsonNewParcel)]))
	}
}

func TestNewShipment(t *testing.T) {
	fromAddress, err := NewAddress(&Address{
		Name:    "Steven Nelson",
		Street1: "1111 Main St",
		Street2: "Suite 204",
		City:    "Conway",
		State:   "AR",
		Zip:     "72032",
		Country: "US",
		Phone:   "5017335509",
		Email:   "stevenrnelson@gmail.com",
	})
	if !HasPrefix(fromAddress.Id, "adr_") || err != nil {
		t.Fatal("creating from address for shipment failed")
	}

	toAddress, err := NewAddress(&Address{
		Name:    "Earth Class Mail",
		Street1: "1538 W 21st St",
		City:    "Houston",
		State:   "TX",
		Zip:     "77008-3642",
		Country: "US",
	})
	if !HasPrefix(toAddress.Id, "adr_") || err != nil {
		t.Fatal("creating to address for shipment failed")
	}

	parcel, err := NewParcel(&Parcel{
		Length: 17.5,
		Width:  12,
		Height: 15.5,
		Weight: 124,
	})
	if !HasPrefix(parcel.Id, "prcl_") || err != nil {
		t.Fatal("creating parcel for shipment failed")
	}

	shipment, err := NewShipment(&Shipment{
		ToAddress:   toAddress,
		FromAddress: fromAddress,
		Parcel:      parcel,
	})
	if !HasPrefix(shipment.Id, "shp_") || err != nil {
		t.Fatal("creating test shipment failed")
	}
}

func TestRetrieveRates(t *testing.T) {
	fromAddress, err := NewAddress(&Address{
		Name:    "Steven Nelson",
		Street1: "1111 Main St",
		Street2: "Suite 204",
		City:    "Conway",
		State:   "AR",
		Zip:     "72032",
		Country: "US",
		Phone:   "5017335509",
		Email:   "stevenrnelson@gmail.com",
	})
	if !HasPrefix(fromAddress.Id, "adr_") || err != nil {
		t.Fatal("creating from address to get rates failed")
	}

	toAddress, err := NewAddress(&Address{
		Name:    "Earth Class Mail",
		Street1: "1538 W 21st St",
		City:    "Houston",
		State:   "TX",
		Zip:     "77008-3642",
		Country: "US",
	})
	if !HasPrefix(toAddress.Id, "adr_") || err != nil {
		t.Fatal("creating to address to get rates failed")
	}

	parcel, err := NewParcel(&Parcel{
		Length: 17.5,
		Width:  12,
		Height: 15.5,
		Weight: 124,
	})
	if !HasPrefix(parcel.Id, "prcl_") || err != nil {
		t.Fatal("creating parcel to get rates failed")
	}

	shipment, err := NewShipment(&Shipment{
		ToAddress:   toAddress,
		FromAddress: fromAddress,
		Parcel:      parcel,
	})
	if !HasPrefix(shipment.Id, "shp_") || err != nil {
		t.Fatal("creating test shipment to get rates failed")
	}
	rates, err := RetrieveRates(shipment.Id)
	if err != nil {
		t.Log(rates)
		t.Fatal("getting rates failed")
	}
}

func TestRetrieveShipment(t *testing.T) {
	fromAddress, err := NewAddress(&Address{
		Name:    "Steven Nelson",
		Street1: "1111 Main St",
		Street2: "Suite 204",
		City:    "Conway",
		State:   "AR",
		Zip:     "72032",
		Country: "US",
		Phone:   "5017335509",
		Email:   "stevenrnelson@gmail.com",
	})
	if !HasPrefix(fromAddress.Id, "adr_") || err != nil {
		t.Fatal("creating from address for shipment retrieval failed")
	}

	toAddress, err := NewAddress(&Address{
		Name:    "Earth Class Mail",
		Street1: "1538 W 21st St",
		City:    "Houston",
		State:   "TX",
		Zip:     "77008-3642",
		Country: "US",
	})
	if !HasPrefix(toAddress.Id, "adr_") || err != nil {
		t.Fatal("creating to address for shipment retrieval failed")
	}

	parcel, err := NewParcel(&Parcel{
		Length: 17.5,
		Width:  12,
		Height: 15.5,
		Weight: 124,
	})
	if !HasPrefix(parcel.Id, "prcl_") || err != nil {
		t.Fatal("creating parcel for shipment retrieval failed")
	}

	shipment, err := NewShipment(&Shipment{
		ToAddress:   toAddress,
		FromAddress: fromAddress,
		Parcel:      parcel,
	})
	if !HasPrefix(shipment.Id, "shp_") || err != nil {
		t.Fatal("creating test shipment retrieval failed")
	}

	newShipment, err := RetrieveShipment(shipment.Id)
	if err != nil {
		t.Log(shipment.Id)
		t.Fatal("couldn't retrieve shipment")
	}
	if fromAddress.Id != newShipment.FromAddress.Id ||
		toAddress.Id != newShipment.ToAddress.Id ||
		parcel.Id != newShipment.Parcel.Id ||
		shipment.Id != newShipment.Id {
		jsonShipment, _ := json.Marshal(shipment)
		jsonNewShipment, _ := json.Marshal(newShipment)
		t.Log(string(jsonShipment[:len(jsonShipment)]))
		t.Log(string(jsonNewShipment[:len(jsonNewShipment)]))
		t.Fatal("Shipments didn't match")
	}
}
func TestBuyShippingLabel(t *testing.T) {
	fromAddress := Address{
		Name:    "Steven Nelson",
		Street1: "1111 Main St",
		Street2: "Suite 204",
		City:    "Conway",
		State:   "AR",
		Zip:     "72032",
		Country: "US",
		Phone:   "5017335509",
		Email:   "stevenrnelson@gmail.com",
	}
	/*
		if !HasPrefix(fromAddress.Id, "adr_") || err != nil {
			t.Fatal("creating from address to buy shipment failed")
		}*/

	toAddress := Address{
		Name:    "Earth Class Mail",
		Street1: "1538 W 21st St",
		City:    "Houston",
		State:   "TX",
		Zip:     "77008-3642",
		Country: "US",
	}
	/*if !HasPrefix(toAddress.Id, "adr_") || err != nil {
		t.Fatal("creating to address to buy shipment failed")
	}*/

	parcel := Parcel{
		Length: 17.5,
		Width:  12,
		Height: 15.5,
		Weight: 124,
	}
	/*if !HasPrefix(parcel.Id, "prcl_") || err != nil {
		t.Fatal("creating parcel to buy shipment failed")
	}*/

	shipment, err := NewShipment(&Shipment{
		ToAddress:   toAddress,
		FromAddress: fromAddress,
		Parcel:      parcel,
	})
	if !HasPrefix(shipment.Id, "shp_") || err != nil {
		t.Fatal("creating test shipment to buy label failed")
	}
	rates, err := RetrieveRates(shipment.Id)
	if err != nil {
		t.Log(rates)
		t.Fatal("getting rates failed")
	}
	label, err := BuyShippingLabel(shipment.Id, rates[0].Id)
	if label.Object != "PostageLabel" || err != nil {
		t.Log(label)
		t.Log(shipment)
		t.Log(shipment.Id)
		t.Log(rates[0])
		t.Log(rates[0].Id)
		t.Log(err)
		t.Fatal("buying shipping label failed")
	}
	fmt.Println("Shipping label bought")
	fmt.Println(label)
}

/*
func TestNewRefund(t *testing.T) {
  fromAddress, err := NewAddress(&Address{
    Name: "Steven Nelson",
    Street1: "1111 Main St",
    Street2: "Suite 204",
    City: "Conway",
    State: "AR",
    Zip: "72032",
    Country: "US",
    Phone: "5017335509",
    Email: "stevenrnelson@gmail.com",
    })
  if !HasPrefix(fromAddress.Id, "adr_") || err != nil{
    t.Fatal("creating from address to buy shipment failed")
  }

  toAddress, err := NewAddress(&Address{
    Name: "Earth Class Mail",
    Street1: "1538 W 21st St",
    City: "Houston",
    State: "TX",
    Zip: "77008-3642",
    Country: "US",
    })
  if !HasPrefix(toAddress.Id, "adr_") || err != nil{
    t.Fatal("creating to address to buy shipment failed")
  }

  parcel, err := NewParcel(&Parcel{
    Length: 17.5,
    Width: 12,
    Height: 15.5,
    Weight: 124,
    })
  if !HasPrefix(parcel.Id, "prcl_") || err != nil {
    t.Fatal("creating parcel to buy shipment failed")
  }

  shipment, err := NewShipment(&Shipment{
    ToAddress: toAddress,
    FromAddress: fromAddress,
    Parcel: parcel,
  })
  if !HasPrefix(shipment.Id, "shp_") || err != nil {
    t.Fatal("creating test shipment to buy label failed")
  }
  rates, err := RetrieveRates(shipment.Id)
  if  err != nil {
    t.Log(rates)
    t.Fatal("getting rates failed")
  }
  label, err := BuyShippingLabel(shipment.Id, rates[0].Id)
  if label.Object != "PostageLabel" || err != nil {
    t.Log(label)
    t.Log(shipment)
    t.Log(shipment.Id)
    t.Log(rates[0])
    t.Log(rates[0].Id)
    t.Log(err)
    t.Fatal("buying shipping label failed")
  }

  refund, err := NewRefund(shipment.Id)
  if refund.Object != "Refund" || err != nil {
    t.Log(shipment.Id)
    t.Log(refund)
    t.Log(refund.Id)
    t.Log(err)
    t.Fatal("creating refund failed")
  }

}
/*
func TestNewCustomsItem(t *testing.T){
  customsItem, err := NewCustomsItem(&CustomsItem{
    Description: "It's bananas",
    Quantity: 144,
    Value: "84",
    Weight: 900,
    HsTariffNumber: "194723",
    OriginCountry: "US",
    })
  if !HasPrefix(customsItem.Id, "cstitem_") || err != nil{
    t.Log(customsItem.Id)
    t.Log(customsItem.Description)
    t.Log(err)
    t.Log(customsItem)
    t.Fatal("creating test customsItem failed")
  }
}

func TestRetrieveCustomsItems(t *testing.T) {
  customsItem, err := NewCustomsItem(&CustomsItem{
    Description: "It's bananas",
    Quantity: 144,
    Value: "84",
    Weight: 900,
    HsTariffNumber: "194723",
    OriginCountry: "US",
    })
  if !HasPrefix(customsItem.Id, "cstitem_") || err != nil{
    t.Log(customsItem.Id)
    t.Log(customsItem.Description)
    t.Log(err)
    t.Log(customsItem)
    t.Fatal("creating test customsItem for retrieval failed")
  }
  newCustomsItem, err := RetrieveCustomsItem(customsItem.Id)
  if err != nil {
    t.Log(newCustomsItem.Id)
    t.Fatal(err)
  }
  if newCustomsItem.Description != customsItem.Description &&
      newCustomsItem.Value != customsItem.Value {
    newcst, _ := json.Marshal(newCustomsItem)
    sentcst, _ := json.Marshal(customsItem)
    t.Log(string(newcst[:len(newcst) ] ))
    t.Log(string(sentcst[:len(sentcst) ] ))
    t.Fatal("customs items don't match")
  }
}

func TestNewCustomsInfo(t *testing.T) {
  customsItem1, err := NewCustomsItem(&CustomsItem{
    Description: "It's bananas",
    Quantity: 144,
    Value: "84",
    Weight: 900,
    HsTariffNumber: "194723",
    OriginCountry: "US",
    })
  customsItem2, err := NewCustomsItem(&CustomsItem{
    Description: "It's apples",
    Quantity: 16,
    Value: "3.50",
    Weight: 25,
    HsTariffNumber: "235163",
    OriginCountry: "US",
    })
  itemSlice := []CustomsItem{customsItem1, customsItem2}
  customsInfo, err := NewCustomsInfo(&CustomsInfo{
    ContentsExplanation: "", //"There's stuff in here",
    ContentsType: "merchandise",
    CustomsCertify: true,
    CustomsSigner: "Franky G",
    EelPfc: "EEL",
    NonDeliveryOption: "return",
    RestrictionComments: "",
    RestrictionType: "none",
    CustomsItems: itemSlice,
    })
    if !HasPrefix(customsInfo.Id, "cstinfo_") || err != nil{
      t.Log(customsInfo.Id)
      t.Log(customsInfo.ContentsExplanation)
      t.Log(err)
      t.Log(customsInfo)
      t.Fatal("creating test customsinfo for retrieval failed")
    }
}

func TestRetrieveCustomsInfo(t *testing.T) {
  customsItem1, err := NewCustomsItem(&CustomsItem{
    Description: "It's bananas",
    Quantity: 144,
    Value: "84",
    Weight: 900,
    HsTariffNumber: "194723",
    OriginCountry: "US",
    })
  customsItem2, err := NewCustomsItem(&CustomsItem{
    Description: "It's apples",
    Quantity: 16,
    Value: "3.50",
    Weight: 25,
    HsTariffNumber: "235163",
    OriginCountry: "US",
    })
  itemSlice := []CustomsItem{customsItem1, customsItem2}
  customsInfo, err := NewCustomsInfo(&CustomsInfo{
    ContentsExplanation: "", //"There's stuff in here",
    ContentsType: "merchandise",
    CustomsCertify: true,
    CustomsSigner: "Franky G",
    EelPfc: "EEL",
    NonDeliveryOption: "return",
    RestrictionComments: "",
    RestrictionType: "none",
    CustomsItems: itemSlice,
    })
  if !HasPrefix(customsInfo.Id, "cstinfo_") || err != nil{
    t.Log(customsInfo.Id)
    t.Log(customsInfo.ContentsExplanation)
    t.Log(err)
    t.Log(customsInfo)
    t.Fatal("creating test customsinfo for retrieval failed")
  }
  newCstInfo, err := RetrieveCustomsInfo(customsInfo.Id)
  if err != nil {
    t.Log(newCstInfo)
    t.Fatal("couldn't retrieve customsinfo")
  }
  if customsInfo.Id != newCstInfo.Id {
    jsonCstInfo, _ := json.Marshal(customsInfo)
    jsonNewCstInfo, _ := json.Marshal(newCstInfo)
    t.Log(string(jsonCstInfo[:len(jsonCstInfo)]))
    t.Log(string(jsonNewCstInfo[:len(jsonNewCstInfo)]))
    t.Fatal("customsinfo didn't match")
  }
}

/*
 * Just to keep the ego in check
func TestMurphy(t *testing.T){
  t.Fatal("Murphy's up to no good")
}
*/

// Copied from strings package
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}
