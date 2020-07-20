package efishery_storage

import(	
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/go-resty/resty/v2"
	"storage-service/domain"
)

func Fetch ()(*[]domain.Storage, error){
	var storages  []domain.Storage
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")

	err = json.Unmarshal(resp.Body(), &storages)
	if err != nil{
		logrus.Error("failed to unmarshal response from rest API storage : " + err.Error())
		return nil, domain.ErrInternalServerError
	}
	return &storages, nil
}