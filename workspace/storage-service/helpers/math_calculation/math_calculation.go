package math

import(
	"storage-service/domain"
)
func CalcMedian(n int) float64 {
	if n%2 == 0 {
		return ((float64(n) / 2) + (float64(n)/2 + 1)) / 2
	} else {
		return (float64(n) + 1) / 2
	}
}

func GetSum(lists []domain.Storage, fn func(i *domain.Storage) float64) float64 {
	var sum float64

	for _, list := range lists {
		sum += fn(&list)
	}

	return sum
}