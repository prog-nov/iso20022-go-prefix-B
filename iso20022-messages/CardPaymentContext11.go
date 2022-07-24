package iso20022

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext11 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext11 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext11) AddPaymentContext() *PaymentContext11 {
	c.PaymentContext = new(PaymentContext11)
	return c.PaymentContext
}

func (c *CardPaymentContext11) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
