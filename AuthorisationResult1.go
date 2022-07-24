package iso20022

// Outcome of the authorisation, and actions to perform.
type AuthorisationResult1 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification33 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Indicates whether the acquirer requires a further exchange completion after the completion of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Instructs the point of interaction (POI) how to contact the host to initiate the maintenance of the terminal.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult1) AddAuthorisationEntity() *GenericIdentification33 {
	a.AuthorisationEntity = new(GenericIdentification33)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult1) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult1) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult1) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *AuthorisationResult1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
