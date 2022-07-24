package iso20022

// Details of the card transaction.
type CardTransactionDetail2 struct {

	// Amounts of the transaction expressed within the terminal currency.
	TransactionAmounts *CardTransactionAmount2 `xml:"TxAmts"`

	// Additional amounts from the processor or the issuer without financial impacts on the transaction amount.
	AdditionalAmounts []*DetailedAmount10 `xml:"AddtlAmts,omitempty"`

	// Account involved in the card transaction.
	AccountAndBalance []*CardAccount2 `xml:"AcctAndBal,omitempty"`

	// Results of the verifications performed by the various agents during the processing of the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	// It corresponds to ISO 8583, field number 57 for the version 93, and 3 for the version 2003.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Data related to an integrated circuit card application.
	// It corresponds to ISO 8583, field number 55 for the versions 93 and 2003.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardTransactionDetail2) AddTransactionAmounts() *CardTransactionAmount2 {
	c.TransactionAmounts = new(CardTransactionAmount2)
	return c.TransactionAmounts
}

func (c *CardTransactionDetail2) AddAdditionalAmounts() *DetailedAmount10 {
	newValue := new(DetailedAmount10)
	c.AdditionalAmounts = append(c.AdditionalAmounts, newValue)
	return newValue
}

func (c *CardTransactionDetail2) AddAccountAndBalance() *CardAccount2 {
	newValue := new(CardAccount2)
	c.AccountAndBalance = append(c.AccountAndBalance, newValue)
	return newValue
}

func (c *CardTransactionDetail2) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardTransactionDetail2) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardTransactionDetail2) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
