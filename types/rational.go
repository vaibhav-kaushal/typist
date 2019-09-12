package types

type Rational struct {
	numerator   int64
	denominator int64
}

func NewRational(numerator int64, denominator int64) *Rational {
	return &Rational{
		denominator: denominator,
		numerator: numerator,
	}
}

