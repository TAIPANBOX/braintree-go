package braintree

type PayPalAccountGateway struct {
	*Braintree
}

func (g *PayPalAccountGateway) Create(paymentMethod *PaymentMethod) (*PayPalAccount, error) {
	resp, err := g.execute("POST", "payment_methods", paymentMethod)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.paypalAccount()
	}
	return nil, &invalidResponseError{resp}
}
