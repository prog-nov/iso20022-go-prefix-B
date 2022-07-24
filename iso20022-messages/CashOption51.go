package iso20022

// Provides information about the cash option.
type CashOption51 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts38 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate47 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms24 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails26 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails22 `xml:"PricDtls,omitempty"`
}

func (c *CashOption51) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption51) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption51) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption51) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption51) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption51) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption51) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption51) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption51) AddAmountDetails() *CorporateActionAmounts38 {
	c.AmountDetails = new(CorporateActionAmounts38)
	return c.AmountDetails
}

func (c *CashOption51) AddDateDetails() *CorporateActionDate47 {
	c.DateDetails = new(CorporateActionDate47)
	return c.DateDetails
}

func (c *CashOption51) AddForeignExchangeDetails() *ForeignExchangeTerms24 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms24)
	return c.ForeignExchangeDetails
}

func (c *CashOption51) AddRateAndAmountDetails() *RateDetails26 {
	c.RateAndAmountDetails = new(RateDetails26)
	return c.RateAndAmountDetails
}

func (c *CashOption51) AddPriceDetails() *PriceDetails22 {
	c.PriceDetails = new(PriceDetails22)
	return c.PriceDetails
}
