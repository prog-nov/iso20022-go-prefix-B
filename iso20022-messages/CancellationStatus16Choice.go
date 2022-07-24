package iso20022

// Specifies whether the status is provided with a reason or not.
type CancellationStatus16Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the cancellation status.
	Reason []*CancellationReason12 `xml:"Rsn"`
}

func (c *CancellationStatus16Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus16Choice) AddReason() *CancellationReason12 {
	newValue := new(CancellationReason12)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
