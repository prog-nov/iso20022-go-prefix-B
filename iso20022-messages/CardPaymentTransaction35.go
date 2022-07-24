package iso20022

// Transaction information in the cancellation response.
type CardPaymentTransaction35 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails6 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction35) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction35) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) AddTransactionDetails() *CardPaymentTransactionDetails6 {
	c.TransactionDetails = new(CardPaymentTransactionDetails6)
	return c.TransactionDetails
}
