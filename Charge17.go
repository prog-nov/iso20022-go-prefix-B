package iso20022

// Amount of money associated with a service.
type Charge17 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType11Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis2Code `xml:"ChrgBsis,omitempty"`

	// Method used to calculate a charge.
	ExtendedChargeBasis *Extended350Code `xml:"XtndedChrgBsis,omitempty"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge17) SetType(value string) {
	c.Type = (*ChargeType11Code)(&value)
}

func (c *Charge17) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge17) SetChargeBasis(value string) {
	c.ChargeBasis = (*TaxationBasis2Code)(&value)
}

func (c *Charge17) SetExtendedChargeBasis(value string) {
	c.ExtendedChargeBasis = (*Extended350Code)(&value)
}

func (c *Charge17) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge17) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge17) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}
