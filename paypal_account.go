package braintree

import "time"

type PayPalAccount struct {
	CustomerId         string         `xml:"customer-id,omitempty"`
	CreatedAt          *time.Time     `xml:"created-at,omitempty"`
	UpdatedAt          *time.Time     `xml:"updated-at,omitempty"`
	Default            bool           `xml:"default,omitempty"`
	Email              string         `xml:"email,omitempty"`
	ImageURL           string         `xml:"image-url,omitempty"`
	Subscriptions      *Subscriptions `xml:"subscriptions,omitempty"`
	Token              string         `xml:"token,omitempty"`
	PaymentMethodNonce string         `xml:"payment-method-nonce,omitempty"`
}

type PayPalAccounts struct {
	PayPalAccount []*PayPalAccount `xml:"paypal-account"`
}
