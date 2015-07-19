package easypost

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var EasyPostApi = map[string]string{
	"Key": "", // Should be specified with a statement like
	// easypost.EasyPostApi["Key"] = "alphabetathetakey"
	// Be sure to do that before calling anything else. Or it will fail
	"BaseUrl": "https://api.easypost.com/v2",
}

func NewAddress(addr *Address) (newAddress Address, err error) {
	data := url.Values{}
	data.Set("address[name]", addr.Name)
	data.Set("address[street1]", addr.Street1)
	data.Set("address[street2]", addr.Street2)
	data.Set("address[city]", addr.City)
	data.Set("address[state]", addr.State)
	data.Set("address[zip]", addr.Zip)
	data.Set("address[country]", addr.Country)
	data.Set("address[phone]", addr.Phone)
	data.Set("address[email]", addr.Email)
	response, err := apiCall("/addresses", data)
	if err == nil {
		err = handleJson(response, &newAddress)
	}
	return newAddress, err
}

func RetrieveAddress(addressId string) (newAddress Address, err error) {
	response, err := apiCall("/addresses/"+addressId, url.Values{})
	if err == nil {
		err = handleJson(response, &newAddress)
	}
	return newAddress, err
}

func (addr *Address) Verify() (message EasyPostMessage, err error) {
	response, err := apiCall("/addresses/"+addr.Id+"/verify", url.Values{})
	var verifiedAddress *VerifiedAddress
	if err == nil {
		err = handleJson(response, &verifiedAddress)
		addr = &verifiedAddress.Address
		message = verifiedAddress.Message
	}
	return message, err
}

func NewParcel(parc *Parcel) (newParcel Parcel, err error) {
	data := url.Values{}
	data.Set("parcel[length]", strconv.FormatFloat(parc.Length, 'f', -1, 64))
	data.Set("parcel[width]", strconv.FormatFloat(parc.Width, 'f', -1, 64))
	data.Set("parcel[height]", strconv.FormatFloat(parc.Height, 'f', -1, 64))
	data.Set("parcel[weight]", strconv.FormatFloat(parc.Weight, 'f', -1, 64))
	response, err := apiCall("/parcels", data)
	if err == nil {

		err = handleJson(response, &newParcel)
	}
	return newParcel, err
}

func RetrieveParcel(parcelId string) (newParcel Parcel, err error) {
	response, err := apiCall("/parcels/"+parcelId, url.Values{})
	if err == nil {
		err = handleJson(response, &newParcel)
	}
	return newParcel, err
}

func NewShipment(shipment *Shipment) (newShipment Shipment, err error) {
	data := url.Values{}
	if len(shipment.ToAddress.Id) > 0 {
		data.Set("shipment[to_address][id]", shipment.ToAddress.Id)
	} else {
		data.Set("shipment[to_address][name]", shipment.ToAddress.Name)
		data.Set("shipment[to_address][company]", shipment.ToAddress.Company)
		data.Set("shipment[to_address][street1]", shipment.ToAddress.Street1)
		data.Set("shipment[to_address][street2]", shipment.ToAddress.Street2)
		data.Set("shipment[to_address][city]", shipment.ToAddress.City)
		data.Set("shipment[to_address][state]", shipment.ToAddress.State)
		data.Set("shipment[to_address][zip]", shipment.ToAddress.Zip)
		data.Set("shipment[to_address][country]", shipment.ToAddress.Country)
		data.Set("shipment[to_address][phone]", shipment.ToAddress.Phone)
		data.Set("shipment[to_address][phoneNumber]", shipment.ToAddress.Phone)
		data.Set("shipment[to_address][email]", shipment.ToAddress.Email)
	}
	if len(shipment.FromAddress.Id) > 0 {
		data.Set("shipment[from_address][id]", shipment.FromAddress.Id)
	} else {
		data.Set("shipment[from_address][name]", shipment.FromAddress.Name)
		data.Set("shipment[from_address][company]", shipment.FromAddress.Company)
		data.Set("shipment[from_address][street1]", shipment.FromAddress.Street1)
		data.Set("shipment[from_address][street2]", shipment.FromAddress.Street2)
		data.Set("shipment[from_address][city]", shipment.FromAddress.City)
		data.Set("shipment[from_address][state]", shipment.FromAddress.State)
		data.Set("shipment[from_address][zip]", shipment.FromAddress.Zip)
		data.Set("shipment[from_address][country]", shipment.FromAddress.Country)
		data.Set("shipment[from_address][phone]", shipment.FromAddress.Phone)
		data.Set("shipment[from_address][phoneNumber]", shipment.FromAddress.Phone)
		data.Set("shipment[from_address][email]", shipment.FromAddress.Email)
	}
	if len(shipment.Parcel.Id) > 0 {
		data.Set("shipment[parcel][id]", shipment.Parcel.Id)
	} else {
		data.Set("shipment[parcel][length]",
			strconv.FormatFloat(shipment.Parcel.Length, 'f', -1, 64))
		data.Set("shipment[parcel][width]",
			strconv.FormatFloat(shipment.Parcel.Width, 'f', -1, 64))
		data.Set("shipment[parcel][height]",
			strconv.FormatFloat(shipment.Parcel.Height, 'f', -1, 64))
		data.Set("shipment[parcel][weight]",
			strconv.FormatFloat(shipment.Parcel.Weight, 'f', -1, 64))
	}
	if shipment.CustomsInfo.Id != "" {
		data.Set("shipment[customs_info][id]", shipment.CustomsInfo.Id)
	}
	response, err := apiCall("/shipments", data)

	if err == nil {
		err = handleJson(response, &newShipment)
	}
	for key, v := range newShipment.Rates {
		newShipment.Rates[key].ServiceName = rateMap[v.Service]
		newShipment.Rates[key].RateFloat, _ = strconv.ParseFloat(v.Rate, 64)
	}
	return newShipment, err
}

func RetrieveShipment(shipmentId string) (newShipment Shipment, err error) {
	response, err := apiCall("/shipments/"+shipmentId, url.Values{})
	if err == nil {
		err = handleJson(response, &newShipment)
	}
	return newShipment, err
}

func RetrieveRates(shipmentId string) (rates []Rate, err error) {
	response, err := apiCall("/shipments/"+shipmentId+"/rates", url.Values{})
	var container Shipment // Dummy shipping object to unmarshal rates into
	if err == nil {
		err = handleJson(response, &container)
	}

	for _, v := range container.Rates {
		v.ServiceName = rateMap[v.Service]
	}
	return container.Rates, err // Return only the rates array.
}

/*
 * Buy shipment
 */
func BuyShippingLabel(shipmentId string, rateId string) (postageLabel PostageLabel,
	err error) {
	data := url.Values{}
	data.Set("rate[id]", rateId)
	response, err := apiCall("/shipments/"+shipmentId+"/buy", data)
	var temp = map[string]json.RawMessage{}

	if err == nil {
		err = handleJson(response, &temp)
		if err != nil {
			fmt.Println("Error parsing response from label purchase" +
				err.Error() + ": " + string(response))
		}
		err = handleJson(temp["postage_label"], &postageLabel)
		if err != nil {
			fmt.Println("Error parsing postage label for " + shipmentId)
			fmt.Println(err.Error() + ": " + string(response))
		}
	}
	return postageLabel, err
}

func NewCustomsItem(customsItem *CustomsItem) (newCustomsItem CustomsItem,
	err error) {
	data := url.Values{}
	data.Set("customs_item[description]", customsItem.Description)
	data.Set("customs_item[quantity]", strconv.FormatFloat(customsItem.Quantity, 'f', -1, 64))
	data.Set("customs_item[value]", customsItem.Value)
	data.Set("customs_item[weight]", strconv.FormatFloat(customsItem.Weight, 'f', -1, 64))
	data.Set("customs_item[hs_tariff_number]", customsItem.HsTariffNumber)
	data.Set("customs_item[origin_country]", customsItem.OriginCountry)
	response, err := apiCall("/customs_items", data)
	if err == nil {
		err = handleJson(response, &newCustomsItem)
	}
	return newCustomsItem, err
}

func RetrieveCustomsItem(customsItemId string) (newCustomsItem CustomsItem,
	err error) {
	response, err := apiCall("/customs_items/"+customsItemId, url.Values{})
	if err == nil {
		err = handleJson(response, &newCustomsItem)
	}
	return newCustomsItem, err
}

func NewCustomsInfo(customsInfo *CustomsInfo) (newCustomsInfo CustomsInfo,
	err error) {
	data := url.Values{}
	data.Set("customs_info[customs_certify]", strconv.FormatBool(customsInfo.CustomsCertify))
	data.Set("customs_info[customs_signer]", customsInfo.CustomsSigner)
	data.Set("customs_info[contents_type]", customsInfo.ContentsType)
	data.Set("customs_info[contents_explanation]", customsInfo.ContentsExplanation)
	data.Set("customs_info[restriction_type]", customsInfo.RestrictionType)
	data.Set("customs_info[eel_pfc]", customsInfo.EelPfc)

	for index, val := range customsInfo.CustomsItems {
		prefix := "customs_info[customs_items][" + strconv.Itoa(index) + "]"
		data.Set(prefix+"[id]", val.Id)
		data.Set(prefix+"[description]", val.Description)
		data.Set(prefix+"[quantity]", strconv.FormatFloat(val.Quantity, 'f', -1, 64))
		// Docs said value is numeric, implementation returns a string
		//data.Set( prefix + "[value]", strconv.FormatFloat(val.Value, 'f', -1, 64))
		data.Set(prefix+"[value]", val.Value)
		data.Set(prefix+"[weight]", strconv.FormatFloat(val.Weight, 'f', -1, 64))
		data.Set(prefix+"[hs_tariff_number]", val.HsTariffNumber)
		data.Set(prefix+"[origin_country]", val.OriginCountry)
	}
	response, err := apiCall("/customs_infos", data)
	if err == nil {
		err = handleJson(response, &newCustomsInfo)
	}
	return newCustomsInfo, err
}

func RetrieveCustomsInfo(customsInfoId string) (newCustomsInfo CustomsInfo,
	err error) {
	response, err := apiCall("/customs_infos/"+customsInfoId, url.Values{})
	if err == nil {
		err = handleJson(response, &newCustomsInfo)
	}
	return newCustomsInfo, err
}

func NewRefund(shipmentId string) (newRefund Refund, err error) {
	response, err := apiCall("/shipments/"+shipmentId+"/refund", url.Values{})
	if err == nil {
		err = handleJson(response, &newRefund)
	}
	return newRefund, err
}

// NewRefundOutsideEasyPost will not likely handle more than one tracking code
// at a time.
func NewRefundOutsideEasyPost(carrier string,
	trackingCodes string) (newRefund Refund, err error) {
	data := url.Values{}
	data.Set("refund[carrier]", carrier)
	data.Set("refund[tracking_codes]", trackingCodes)
	response, err := apiCall("/refunds", data)
	if err == nil {
		err = handleJson(response, &newRefund)
	}
	return newRefund, err
}

func RetrieveRefund(refundId string) (newRefund Refund, err error) {
	response, err := apiCall("/refunds/"+refundId, url.Values{})
	if err == nil {
		err = handleJson(response, &newRefund)
	}
	return newRefund, err
}

func NewBatch(shipments []Shipment, createAndBuy bool) (newBatch Batch, err error) {
	data := url.Values{}

	for index, val := range shipments {
		prefix := "batch[shipment][" + strconv.Itoa(index) + "]"
		data.Set(prefix+"[from_address][name]", val.FromAddress.Name)
		data.Set(prefix+"[from_address][company]", val.FromAddress.Company)
		data.Set(prefix+"[from_address][street1]", val.FromAddress.Street1)
		data.Set(prefix+"[from_address][street2]", val.FromAddress.Street2)
		data.Set(prefix+"[from_address][city]", val.FromAddress.City)
		data.Set(prefix+"[from_address][state]", val.FromAddress.State)
		data.Set(prefix+"[from_address][zip]", val.FromAddress.Zip)
		data.Set(prefix+"[from_address][country]", val.FromAddress.Country)
		data.Set(prefix+"[from_address][phone]", val.FromAddress.Phone)
		data.Set(prefix+"[from_address][email]", val.FromAddress.Email)
		data.Set(prefix+"[to_address][name]", val.ToAddress.Name)
		data.Set(prefix+"[to_address][company]", val.ToAddress.Company)
		data.Set(prefix+"[to_address][street1]", val.ToAddress.Street1)
		data.Set(prefix+"[to_address][street2]", val.ToAddress.Street2)
		data.Set(prefix+"[to_address][city]", val.ToAddress.City)
		data.Set(prefix+"[to_address][state]", val.ToAddress.State)
		data.Set(prefix+"[to_address][zip]", val.ToAddress.Zip)
		data.Set(prefix+"[to_address][country]", val.ToAddress.Country)
		data.Set(prefix+"[to_address][phone]", val.ToAddress.Phone)
		data.Set(prefix+"[to_address][email]", val.ToAddress.Email)
		data.Set(prefix+"[parcel][length]",
			strconv.FormatFloat(val.Parcel.Length, 'f', -1, 64))
		data.Set(prefix+"[parcel][width]",
			strconv.FormatFloat(val.Parcel.Width, 'f', -1, 64))
		data.Set(prefix+"[parcel][height]",
			strconv.FormatFloat(val.Parcel.Height, 'f', -1, 64))
		data.Set(prefix+"[parcel][weight]",
			strconv.FormatFloat(val.Parcel.Weight, 'f', -1, 64))
		data.Set(prefix+"[reference]", val.Reference)
		if createAndBuy {
			data.Set(prefix+"[carrier]", val.Rates[0].Carrier)
			data.Set(prefix+"[service]", val.Rates[0].Service)
		}
	}
	var response []byte
	if createAndBuy {
		response, err = apiCall("/batches/create_and_buy", data)
	} else {
		response, err = apiCall("/batches", data)
	}

	if err == nil {
		err = handleJson(response, &newBatch)
	}
	return newBatch, err
}

// RetrieveBatchLabel requests the url of a batch label once all postage for
// the batch has been purchased. The label url will not be available until all
// the shipments are in the "postage_purchased" status. labelType can be one of
// two types: "pdf" or "epl2"
func RetreiveBatchLabel(batchId string, labelType string) (
	newBatch Batch, err error) {
	data := url.Values{}
	data.Set("file_format", labelType)
	response, err := apiCall("/batches/"+batchId+"/label", data)

	if err == nil {
		err = handleJson(response, &newBatch)
	}
	return newBatch, err
}

func AddShipmentsToBatch(batchId string, shipmentIds []string) (
	newBatch Batch, err error) {
	return addOrRemoveShipmentsToBatch(batchId, shipmentIds, false)
}

func RemoveShipmentsFromBatch(batchId string, shipmentIds []string) (
	newBatch Batch, err error) {
	return addOrRemoveShipmentsToBatch(batchId, shipmentIds, true)
}

func addOrRemoveShipmentsToBatch(batchId string, shipmentIds []string,
	removeShipment bool) (newBatch Batch, err error) {

	data := url.Values{}
	for index, val := range shipmentIds {
		data.Set("shipments["+strconv.Itoa(index)+"][id]", val)
	}
	var response []byte
	if removeShipment {
		response, err = apiCall("/batches/"+batchId+"remove_shipments", data)
	} else {
		response, err = apiCall("/batches/"+batchId+"add_shipments", data)
	}
	if err == nil {
		err = handleJson(response, &newBatch)
	}
	return newBatch, err
}

func NewScanForm(scanForm *ScanForm) (newScanForm ScanForm, err error) {
	trackingCodes := ""
	for _, val := range scanForm.TrackingCodes {
		trackingCodes += val + ","
	}
	trackingCodes = trackingCodes[:len(trackingCodes)-1]

	data := url.Values{}
	data.Set("scan_form[from_address][name]", scanForm.Address.Name)
	data.Set("scan_form[from_address][company]", scanForm.Address.Company)
	data.Set("scan_form[from_address][street1]", scanForm.Address.Street1)
	data.Set("scan_form[from_address][street2]", scanForm.Address.Street2)
	data.Set("scan_form[from_address][city]", scanForm.Address.City)
	data.Set("scan_form[from_address][state]", scanForm.Address.State)
	data.Set("scan_form[from_address][country]", scanForm.Address.Country)
	data.Set("scan_form[from_address][zip]", scanForm.Address.Zip)
	data.Set("scan_form[from_address][phone]", scanForm.Address.Phone)
	data.Set("scan_form[from_address][email]", scanForm.Address.Email)
	data.Set("scan_form[tracking_codes]", trackingCodes)

	response, err := apiCall("/scan_forms", data)
	if err == nil {
		err = handleJson(response, &newScanForm)
	}
	return newScanForm, err
	err = handleJson(response, &newScanForm)
	return newScanForm, err
}

func RetrieveScanForm(scanFormId string) (newScanForm ScanForm, err error) {
	response, err := apiCall("/scanForm/"+scanFormId, url.Values{})
	if err == nil {
		err = handleJson(response, &newScanForm)
	}
	return newScanForm, err
}

// handleJson provides a thin wrapper around the json.Unmarshal func to catch
// errors returned by the EasyPost API. If an error is returned from the API,
// it is converted into a proper Go error, and returned.
func handleJson(response []byte, target interface{}) error {
	var err error
	/*if 0 != bytes.Compare([]byte("{\"error\":}"), response[:9]) {
	    err = json.Unmarshal(response, &target)
	  } else {
	    var rawResponse interface{}
	    err = json.Unmarshal(response, &rawResponse)
	    parsed := rawResponse.(map[string]interface{})
		  return errors.New(parsed["error"].(string))
	  }*/
	err = json.Unmarshal(response, &target)
	if err != nil {
		fmt.Println(err.Error() + ":  " + string(response))
	}
	epm, ok := target.(EasyPostResponse)
	if err == nil && ok && epm.Error != "" {
		err = errors.New(epm.Error)
	}
	return err
}

func apiCall(path string, data url.Values) (response []byte, err error) {
	if EasyPostApi["Key"] == "" {
		return nil, errors.New("please specify an API key")
	}
	/*
	 * Construct Message Body
	 */
	var postBody bytes.Buffer
	for key, val := range data {
		postBody.WriteString(key + "=" + url.QueryEscape(val[0]) + "&")
	}
	requestMethod := "POST"
	if len(data) > 0 {
		postBody.Truncate(postBody.Len() - 1)
	} else {
		requestMethod = "GET"
	}
	mBody := postBody.Bytes()
	endpointUrl := EasyPostApi["BaseUrl"] + path
	request, err := http.NewRequest(requestMethod, endpointUrl,
		bytes.NewReader(mBody))

	/*
	 * Set Header information
	 */
	// ua := "{'client_version' : '2.0.0', 'lang' : 'go',
	// 'publisher' : '@stevennelson',' 'request_lib': 'net/http'}"
	ua := "Go-EasyPost 2.0.0"
	request.Header.Add("X-EasyPost-Client-User-Agent", ua)
	request.SetBasicAuth(EasyPostApi["Key"], "")

	/*
	 * Create http client, send request, and receive response
	 * In the future, only create a single Client and reuse it
	 */
	client := &http.Client{}
	httpResponse, httpErr := client.Do(request)
	var result []byte

	/*
	 * Handle (or don't) any errors that occured in the transport
	 */
	if httpErr != nil {
		fmt.Println(postBody.Bytes())
		fmt.Println(httpErr.Error())
		return nil, httpErr
	} else if httpResponse.Body == nil {
		fmt.Println(postBody.Bytes())
		return nil, errors.New("response from api is empty")
	} else {
		responseBuffer := new(bytes.Buffer)
		responseBuffer.ReadFrom(httpResponse.Body)
		result = responseBuffer.Bytes()
	}
	if httpResponse.Body != nil {
		httpResponse.Body.Close()
	}
	return result, nil
}
