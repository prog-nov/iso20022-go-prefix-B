package iso20022

// Transaction information in the authorisation response.
type CardPaymentTransaction2 struct {

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails2 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction2) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction2) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction2) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction2) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction2) AddTransactionDetails() *CardPaymentTransactionDetails2 {
	c.TransactionDetails = new(CardPaymentTransactionDetails2)
	return c.TransactionDetails
}
