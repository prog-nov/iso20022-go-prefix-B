package iso20022

// Environment of the card payment transaction.
type CardPaymentEnvironment40 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard11 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder8 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment40) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment40) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment40) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment40) AddCard() *PaymentCard11 {
	c.Card = new(PaymentCard11)
	return c.Card
}

func (c *CardPaymentEnvironment40) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment40) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment40) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment40) AddCardholder() *Cardholder8 {
	c.Cardholder = new(Cardholder8)
	return c.Cardholder
}

func (c *CardPaymentEnvironment40) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}
