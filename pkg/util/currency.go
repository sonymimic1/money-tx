package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	TWD = "TWD"
)

func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, TWD:
		return true
	}
	return false
}
