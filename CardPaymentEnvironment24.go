package iso20022

// Environment of the transaction.
type CardPaymentEnvironment24 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard8 `xml:"Card"`
}

func (c *CardPaymentEnvironment24) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment24) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment24) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment24) AddCard() *PaymentCard8 {
	c.Card = new(PaymentCard8)
	return c.Card
}
