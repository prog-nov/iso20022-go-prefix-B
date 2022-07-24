package iso20022

// Details of the transaction in the authorisation request.
type CardPaymentTransactionDetails19 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails19) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails19) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails19) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails19) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails19) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails19) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails19) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails19) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardPaymentTransactionDetails19) AddCurrencyConversion() *CurrencyConversion3 {
	c.CurrencyConversion = new(CurrencyConversion3)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails19) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails19) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails19) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails19) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails19) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
