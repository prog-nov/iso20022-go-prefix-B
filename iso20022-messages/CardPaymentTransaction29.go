package iso20022

// Transaction information in the batch capture.
type CardPaymentTransaction29 struct {

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails15 `xml:"TxDtls"`

	// Outcome of the authorisation request, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction29) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction29) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction29) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction29) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction29) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction29) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction29) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction29) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction29) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction29) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction29) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) AddTransactionDetails() *CardPaymentTransactionDetails15 {
	c.TransactionDetails = new(CardPaymentTransactionDetails15)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction29) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction29) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction29) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}
