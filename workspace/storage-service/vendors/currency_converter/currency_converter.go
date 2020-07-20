package currency_converter

import(
	"encoding/json"
	"os"
	"net/http"
	"storage-service/domain"
	"github.com/sirupsen/logrus"
	"github.com/go-resty/resty/v2"
)

func GetCurrency()(*domain.Currency, error){
	var(
		currency domain.Currency
		respError domain.ResponseErrorAPI
	)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey="+os.Getenv("CURRENCY_CONVERTER_API_KEY"))
	if resp.StatusCode() == http.StatusBadRequest{
		err := json.Unmarshal(resp.Body(), &respError)
		if err != nil{
			logrus.Error("failed to unmarshal response from rest API currency : " + err.Error())
			return nil, domain.ErrInternalServerError
		}
		logrus.Error(respError.Error)
		return nil, domain.ErrInternalServerError
	}
	err = json.Unmarshal(resp.Body(), &currency)
	if err != nil{
		logrus.Error("failed to unmarshal response from rest API currency : " + err.Error())
		return nil, domain.ErrInternalServerError
	}
	return &currency, nil
}