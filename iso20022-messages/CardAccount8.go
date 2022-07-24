package iso20022

// Customer account information.
type CardAccount8 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType3Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	// Balance of the account.
	Balance *AmountAndDirection43 `xml:"Bal,omitempty"`

	// Indicates if the balance must be displayed to the customer on the ATM.
	BalanceDisplayFlag *TrueFalseIndicator `xml:"BalDispFlg,omitempty"`

	// Indicates if this is the default account.
	DefaultAccountIndicator *TrueFalseIndicator `xml:"DfltAcctInd,omitempty"`
}

func (c *CardAccount8) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount8) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount8) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount8) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount8) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount8) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount8) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount8) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount8) AddBalance() *AmountAndDirection43 {
	c.Balance = new(AmountAndDirection43)
	return c.Balance
}

func (c *CardAccount8) SetBalanceDisplayFlag(value string) {
	c.BalanceDisplayFlag = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount8) SetDefaultAccountIndicator(value string) {
	c.DefaultAccountIndicator = (*TrueFalseIndicator)(&value)
}
