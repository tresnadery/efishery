package mapper

import(
	"fmt"
	"strings"
	"time"
	"sort"
	"strconv"
	"storage-service/domain"
	mathHelper "storage-service/helpers/math_calculation"
)

func Grouping(lists []domain.Storage) map[string][]domain.Storage{
	grouped := make(map[string][]domain.Storage)

	for _, list := range lists {
		if list.ProvinceArea == "" || list.Timestamp == "" || list.Price == "" {
			continue
		}

		i, _ := strconv.Atoi(list.Timestamp)
		t := time.Unix(int64(i), 0).UTC()
		y, w := t.ISOWeek()
		key := fmt.Sprintf("%s#%d#%d", list.ProvinceArea, y, w)
		grouped[key] = append(grouped[key], list)
	}

	return grouped
}

func NewListAggregate(province ,year, week, min, max, med ,avg string) domain.AggregateStorage {
	return domain.AggregateStorage{
		ProvinceArea: province,
		Year:        year,
		Week:     	week,
		Min:          min,
		Max:          max,
		Median:       med,
		Avg:          avg,
	}
}

func ListAggregateStoreage(grouped map[string][]domain.Storage) []domain.AggregateStorage {
	var result []domain.AggregateStorage

	for k, lists := range grouped {
		keys := strings.Split(k, "#")
		sort.SliceStable(lists, func(i, j int) bool {
			priceI, _ := strconv.ParseFloat(lists[i].Price, 64)
			priceJ, _ := strconv.ParseFloat(lists[j].Price, 64)

			return priceI < priceJ
		})
		listLen := len(lists)
		min := lists[0].Price
		max := lists[listLen-1].Price
		med := mathHelper.CalcMedian(len(lists))
		sumOfPrice := mathHelper.GetSum(lists, func(i *domain.Storage) float64 {
			price, _ := strconv.ParseFloat(i.Price, 64)
			return price
		})
		avg := sumOfPrice / float64(listLen)

		result = append(result, NewListAggregate(keys[0], keys[1], keys[2], min, max, fmt.Sprintf("%f", med),
			fmt.Sprintf("%f", avg)))
	}

	return result
}