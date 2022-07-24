package iso20022

// Account to or from which a cash entry is made.
type CashAccount26 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName3 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification2Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, ISA.
	InvestmentAccountType *InvestmentAccountType1Choice `xml:"InvstmtAcctTp,omitempty"`

	// Other identification such as national registration identification number, passport number.
	AccountOwnerOtherIdentification []*GenericIdentification46 `xml:"AcctOwnrOthrId,omitempty"`
}

func (c *CashAccount26) AddIdentification() *AccountIdentificationAndName3 {
	c.Identification = new(AccountIdentificationAndName3)
	return c.Identification
}

func (c *CashAccount26) AddAccountOwner() *PartyIdentification2Choice {
	c.AccountOwner = new(PartyIdentification2Choice)
	return c.AccountOwner
}

func (c *CashAccount26) AddAccountServicer() *PartyIdentification2Choice {
	c.AccountServicer = new(PartyIdentification2Choice)
	return c.AccountServicer
}

func (c *CashAccount26) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}

func (c *CashAccount26) AddInvestmentAccountType() *InvestmentAccountType1Choice {
	c.InvestmentAccountType = new(InvestmentAccountType1Choice)
	return c.InvestmentAccountType
}

func (c *CashAccount26) AddAccountOwnerOtherIdentification() *GenericIdentification46 {
	newValue := new(GenericIdentification46)
	c.AccountOwnerOtherIdentification = append(c.AccountOwnerOtherIdentification, newValue)
	return newValue
}