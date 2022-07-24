package iso20022

// Details of the transaction in the authorisation response contained in a batch.
type CardPaymentTransactionDetails17 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails17) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails17) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails17) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails17) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails17) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails17) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
