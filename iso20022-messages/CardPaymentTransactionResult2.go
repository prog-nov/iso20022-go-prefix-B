package iso20022

// Result of the transaction.
type CardPaymentTransactionResult2 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification70 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation from the acquirer.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (c *CardPaymentTransactionResult2) AddAuthorisationEntity() *GenericIdentification70 {
	c.AuthorisationEntity = new(GenericIdentification70)
	return c.AuthorisationEntity
}

func (c *CardPaymentTransactionResult2) AddResponseToAuthorisation() *ResponseType1 {
	c.ResponseToAuthorisation = new(ResponseType1)
	return c.ResponseToAuthorisation
}

func (c *CardPaymentTransactionResult2) SetAuthorisationCode(value string) {
	c.AuthorisationCode = (*Min6Max8Text)(&value)
}
