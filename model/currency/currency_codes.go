package currency

type Code string

const (
	USD Code = "USD"
	EUR Code = "EUR"
	CNY Code = "CNY"
	BYN Code = "BYN"
	INR Code = "INR"
	RUB Code = "RUB"
)

func (code Code) String() string {
	return string(code)
}

// DefaultCodes returns currency codes for UI menu in default order
func DefaultCodes() []Code {
	return []Code{
		USD,
		EUR,
		CNY,
		BYN,
		INR,
	}
}
