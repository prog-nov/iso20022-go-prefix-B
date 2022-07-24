package iso20022

// Specifies reasons for the cancelled status.
type CancelledStatusReason12 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason9Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason12) AddReasonCode() *CancelledReason9Choice {
	c.ReasonCode = new(CancelledReason9Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason12) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
