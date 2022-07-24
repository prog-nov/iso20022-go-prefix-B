package iso20022

// Amount of money associated with a service.
type ChargesDetails2 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType8Code `xml:"Tp"`

	// Specifies the type of charge by means of a code.
	OtherChargesType *Max35Text `xml:"OthrChrgsTp"`

	// Amount of money asked or paid for the charge.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (c *ChargesDetails2) SetType(value string) {
	c.Type = (*ChargeType8Code)(&value)
}

func (c *ChargesDetails2) SetOtherChargesType(value string) {
	c.OtherChargesType = (*Max35Text)(&value)
}

func (c *ChargesDetails2) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}
