package braintree

type PaymentMethod struct {
	XMLName            string             `xml:"payment-method"`
	BillingAddress     *Address           `xml:"billing-address,omitempty"`
	BillingAddressId   string             `xml:"billing-address-id,omitempty"`
	CardHolderName     string             `xml:"cardholder-name,omitempty"`
	CustomerId         string             `xml:"customer-id,omitempty"`
	DeviceData         string             `xml:"device-data,omitempty"`
	Options            *CreditCardOptions `xml:"options,omitempty"`
	PaymentMethodNonce string             `xml:"payment-method-nonce,omitempty"`
	Token              string             `xml:"token,omitempty"`
}
