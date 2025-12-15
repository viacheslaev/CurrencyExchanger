package model

type CurrencyCode string

const (
	USD CurrencyCode = "USD"
	EUR CurrencyCode = "EUR"
	CNY CurrencyCode = "CNY"
	BYN CurrencyCode = "BYN"
	INR CurrencyCode = "INR"
	RUB CurrencyCode = "RUB"
)

func (code CurrencyCode) String() string {
	return string(code)
}

// DefaultCodes returns currency codes for UI menu in default order
func DefaultCodes() []CurrencyCode {
	return []CurrencyCode{
		USD,
		EUR,
		CNY,
		BYN,
		INR,
	}
}
