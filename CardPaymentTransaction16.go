package iso20022

// Transaction information in the cancellation advice.
type CardPaymentTransaction16 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason2Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails7 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction16) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction16) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction16) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction16) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction16) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction16) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason2Code)(&value))
}

func (c *CardPaymentTransaction16) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction16) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction16) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction16) AddTransactionDetails() *CardPaymentTransactionDetails7 {
	c.TransactionDetails = new(CardPaymentTransactionDetails7)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction16) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction16) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}
