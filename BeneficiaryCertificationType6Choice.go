package iso20022

// Choice between a standard code or proprietary code to specify the type of beneficiary certification required.
type BeneficiaryCertificationType6Choice struct {

	// Standard code to specify the type of certification required.
	Code *BeneficiaryCertificationType5Code `xml:"Cd"`

	// Proprietary identification of the type of certification required.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType6Choice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType5Code)(&value)
}

func (b *BeneficiaryCertificationType6Choice) AddProprietary() *GenericIdentification20 {
	b.Proprietary = new(GenericIdentification20)
	return b.Proprietary
}
