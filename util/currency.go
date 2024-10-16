package util

const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
	PLN = "PLN"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, JPY, PLN:
		return true
	default:
		return false
	}
}
