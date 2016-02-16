package braintree

import (
	"encoding/xml"
	"github.com/lionelbarrow/braintree-go/nullable"
)

type Customer struct {
	XMLName      string       `xml:"customer"`
	Id           string       `xml:"id,omitempty"`
	FirstName    string       `xml:"first-name,omitempty"`
	LastName     string       `xml:"last-name,omitempty"`
	Company      string       `xml:"company,omitempty"`
	Email        string       `xml:"email,omitempty"`
	Phone        string       `xml:"phone,omitempty"`
	Fax          string       `xml:"fax,omitempty"`
	Website      string       `xml:"website,omitempty"`
	CreditCard   *CreditCard  `xml:"credit-card,omitempty"`
	CreditCards  *CreditCards `xml:"credit-cards,omitempty"`
	CustomFields CustomFields `xml:"custom-fields"`
}

// DefaultCreditCard returns the default credit card, or nil
func (c *Customer) DefaultCreditCard() *CreditCard {
	for _, card := range c.CreditCards.CreditCard {
		if card.Default {
			return card
		}
	}
	return nil
}

type CustomerSearchResult struct {
	XMLName           string              `xml:"customers"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number"`
	PageSize          *nullable.NullInt64 `xml:"page-size"`
	TotalItems        *nullable.NullInt64 `xml:"total-items"`
	Customers         []*Customer         `xml:"customer"`
}

type CustomField struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type CustomFields struct {
	Items []CustomField `xml:",any"`
}

func NewCustomFields(fields map[string]string) CustomFields {
	cfs := CustomFields{}
	for field, value := range fields {
		cfs.Items = append(cfs.Items, CustomField{
			XMLName: xml.Name{
				Local: field,
			},
			Value: value,
		})
	}
	return cfs
}

func (cfs *CustomFields) Map() map[string]string {
	m := make(map[string]string)
	for _, field := range cfs.Items {
		m[field.XMLName.Local] = field.Value
	}
	return m
}

func (cfs *CustomFields) Get(name string) string {
	return cfs.Map()[name]
}
