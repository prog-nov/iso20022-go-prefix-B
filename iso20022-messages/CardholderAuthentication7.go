package iso20022

// Data related to the authentication of the cardholder.
type CardholderAuthentication7 struct {

	// Method and data intended to be used for this transaction to authenticate the cardholder or its card.
	AuthenticationMethod *AuthenticationMethod5Code `xml:"AuthntcnMtd"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN4 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identification of the cardholder to verify.
	CardholderIdentification *PersonIdentification7 `xml:"CrdhldrId,omitempty"`

	// Numeric characters of the cardholder's billing or shipping address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`
}

func (c *CardholderAuthentication7) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod5Code)(&value)
}

func (c *CardholderAuthentication7) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication7) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication7) AddCardholderOnLinePIN() *OnLinePIN4 {
	c.CardholderOnLinePIN = new(OnLinePIN4)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication7) AddCardholderIdentification() *PersonIdentification7 {
	c.CardholderIdentification = new(PersonIdentification7)
	return c.CardholderIdentification
}

func (c *CardholderAuthentication7) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}
