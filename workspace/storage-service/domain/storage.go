package domain

type Storage struct{
	UUID string `json:"uuid"`
	Commodity string `json:"komoditas"`
	ProvinceArea string `json:"area_provinsi"`
	CityArea string `json:"area_kota"`
	Size string `json:"size"`
	Price string `json:"price"`
	PriceInUSD string `json:"price_in_usd"`
	DateParsed string `json:"tgl_parsed"`
	Timestamp string `json:"timestamp"`
}

type AggregateStorage struct{
	ProvinceArea string `json:"province_area"`
	Year string `json:"year"`
	Week string `json:"week"`
	Min  string `json:"min"`
	Max  string `json:"max"`
	Median string `json:"median"`
	Avg string `json:"avg"`
}
type StorageUsecase interface{
	Store(rate float64)(error)
	GetRate()(float64, error)
}

type StorageRepository interface{
	Store(rate float64)(error)
	GetRate()(float64, error)
}