package easypost

import "time"

type EasyPostMessage struct {
	Message string
}

type EasyPostResponse struct {
	Error string
}

type VerifiedAddress struct {
	Address Address
	Message EasyPostMessage
}

type Address struct {
	Id        string
	Object    string
	Error     string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string
	Company   string
	Street1   string
	Street2   string
	City      string
	State     string
	Zip       string
	Country   string
	Email     string
	Phone     string
}

type Rate struct {
	Id          string
	Object      string
	Error       string
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Service     string
	ServiceName string
	Rate        string
	RateFloat   float64
	Carrier     string
	ShipmentId  string `json:"shipment_id"`
}

type ScanForm struct {
	Id            string
	Object        string
	Error         string
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Address       Address
	TrackingCodes []string `json:"tracking_codes"`
	FormUrl       string   `json:"form_url"`
	FormFileType  string   `json:"form_file_type"`
}

type CustomsInfo struct {
	Id                  string
	Object              string
	Error               string
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	ContentsExplanation string        `json:""`
	ContentsType        string        `json:"contents_type"`
	CustomsCertify      bool          `json:"customs_certify"`
	CustomsSigner       string        `json:"customs_signer"`
	EelPfc              string        `json:"eel_pfc"`
	NonDeliveryOption   string        `json:"non_delivery_option"`
	RestrictionComments string        `json:"restriction_comments"`
	RestrictionType     string        `json:"restriction_type"`
	CustomsItems        []CustomsItem `json:"customs_items"`
}

type CustomsItem struct {
	Id             string
	Object         string
	Error          string
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Description    string
	HsTariffNumber string `json:"hs_tariff_number"`
	OriginCountry  string `json:"origin_country"`
	Quantity       float64
	Value          string
	Weight         float64
}

type PostageLabel struct {
	Id              string
	Object          string
	Error           string
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DateAdvance     int64     `json:"date_advance"`
	IntegratedForm  string    `json:"integrated_form"`
	LabelDate       time.Time `json:"label_date"`
	LabelResolution int64     `json:"label_resolution"`
	LabelSize       string    `json:"label_size"`
	LabelType       string    `json:"label_type"`
	LabelFileType   string    `json:"label_file_type"`
	LabelUrl        string    `json:"label_url"`
	LabelPDFUrl     string    `json:"label_pdf_url"`
	LabelEpl2Url    string    `json:"label_epl2_url"`
	LabelZp1Url     string    `json:"label_zp1_url"`
	SelectedRate    Rate      `json:"selected_rate"`
}

type Parcel struct {
	Id                string
	Object            string
	Error             string
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Length            float64
	Width             float64
	Height            float64
	PredefinedPackage string `json:"predefined_package"`
	Weight            float64
}

type Shipment struct {
	Id           string
	Object       string
	Error        string
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Type         string
	ToAddress    Address `json:"to_address"`
	FromAddress  Address `json:"from_address"`
	Parcel       Parcel
	CustomsInfo  CustomsInfo `json:"customs_info"`
	ScanForm     ScanForm    `json:"scan_form"`
	Rates        []Rate
	SelectedRate Rate         `json:"selected_rate"`
	PostageLabel PostageLabel `json:"postage_label"`
	TrackingCode string       `json:"tracking_code"`
	Reference    string
	RefundStatus string `json:"refund_status"`
	Insurance    float64
	BatchStatus  string `json:"batch_status"`
	BatchMessage string `json:"batch_message"`
	Options      ShippingOptions
}

type ShippingOptions struct {
	AddressValidationLevel string `json:"address_validation_level"`
	ByDrone                string `json:"by_drone"`
	Currency               string
	CarbonNeutral          string `json:"carbon_neutral"`
	DateAdvance            string `json:"date_advance"`
	DeclaredValue          string `json:"declared_value"`
	DeliveryConfirmation   string `json:"delivery_confirmation"`
	DryIce                 string `json:"dry_ice"`
	DryIceMedical          string `json:"dry_ice_medical"`
	DryIceWeight           string `json:"dry_ice_weight"`
	InvoiceNumber          string `json:"invoice_number"`
	Machinable             string
	PoFacility             string `json:"po_facility"`
	PoZip                  string `json:"po_zip"`
	PrintCustom1           string `json:"print_custom_1"`
	PrintCustom2           string `json:"print_custom_2"`
	PrintCustom3           string `json:"print_custom_3"`
	ResidentialToAddress   string `json:"residential_to_address"`
	SaturdayDelivery       string `json:"saturday_delivery"`
	SmartPostHub           string `json:"smartpost_hub"`
	SmartPostManifest      string `json:"smartpost_manifest"`
}

type Batch struct {
	Id        string
	Object    string
	Error     string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Shipments []Shipment
	LabelUrl  string `json:"label_url"`
	Status    BatchStatus
}

type BatchStatus struct {
	Created                int64
	PostagePurchased       int64 `json:"postage_purchased"`
	PostagePurchasedFailed int64 `json:"postage_purchase_failed"`
}

type Refund struct {
	Id           string
	Object       string
	Error        string
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	TrackingCode string    `json:"tracking_code"`
	Status       string
	Carrier      string
	ShipmentId   string `json:"shipment_id"`
}
