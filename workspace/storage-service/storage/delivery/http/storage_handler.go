package http

import(
	"fmt"
	"strconv"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"	
	"storage-service/domain"
	authHelper "storage-service/helpers/auth"
	mapperHelper "storage-service/helpers/mapper"
	storageVendor "storage-service/vendors/efishery_storage"
	currencyVendor "storage-service/vendors/currency_converter"
	// validatorHelper "storage-service/helpers/validator"
)

type ResponseError struct{
	Errors interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
}
type ResponseProfile struct{
	Message string `json:"message"`
	User interface{} `json:"user"`
}
type ResponseStorage struct{
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
type StorageHandler struct{
	Conn *sql.DB
	storageUsecase domain.StorageUsecase
}

func NewStorageHandler(ctx *gin.Engine, us domain.StorageUsecase){
	handler := &StorageHandler{
		storageUsecase: us,		
	}
	ctx.GET("storages", authHelper.Middleware("all"), handler.Fetch)
	ctx.GET("admin/storages", authHelper.Middleware("admin"), handler.FetchWithAggregate)
	ctx.GET("users/profiles", handler.Profile)
}
func appendPriceInUSDToStorage(payload *[]domain.Storage, rate float64)(*[]domain.Storage, error){
	var storages []domain.Storage
	for _, val := range *payload{
		convPrice, err := strconv.Atoi(val.Price)
		if err != nil{
			logrus.Error("failed to convert price : " + err.Error())
			return nil, domain.ErrInternalServerError
		}
		storage := val
		storage.PriceInUSD = fmt.Sprintf("%.2f", (float64(convPrice) * rate))
		storages = append(storages, storage)
	}
	return &storages, nil
}

func (s *StorageHandler) Fetch(ctx *gin.Context){
	rate, err := s.storageUsecase.GetRate()
	if err != nil && getStatusCode(err) == http.StatusInternalServerError{
		ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	if getStatusCode(err) == http.StatusNotFound{
		currency, err := currencyVendor.GetCurrency()
		if err != nil{
			ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		if err := s.storageUsecase.Store(currency.IDR_USD); err != nil{
			ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			return	
		}
		rate = currency.IDR_USD
	}
	storages, err := storageVendor.Fetch()
	if err != nil{
		ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	res, err := appendPriceInUSDToStorage(storages, rate)
	if err != nil{
		ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return	
	}
	ctx.JSON(http.StatusOK, ResponseStorage{
		Message: domain.SuccessGetStorage,
		Data: res,
	})
}
func (s *StorageHandler) FetchWithAggregate(ctx *gin.Context){
	rate, err := s.storageUsecase.GetRate()
	if err != nil && getStatusCode(err) == http.StatusInternalServerError{
		ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	if getStatusCode(err) == http.StatusNotFound{
		currency, err := currencyVendor.GetCurrency()
		if err != nil{
			ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
		if err := s.storageUsecase.Store(currency.IDR_USD); err != nil{
			ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			return	
		}
		rate = currency.IDR_USD
	}
	storages, err := storageVendor.Fetch()
	if err != nil{
		ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	res, err := appendPriceInUSDToStorage(storages, rate)
	if err != nil{
		ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return	
	}
	ctx.JSON(http.StatusOK, ResponseStorage{
		Message: domain.SuccessGetStorage,
		Data: mapperHelper.ListAggregateStoreage(mapperHelper.Grouping(*res)),
	})
}
func (s *StorageHandler) Profile(ctx *gin.Context){
	token := ctx.Request.Header.Get("Authorization")
	user, err := authHelper.GetClaimsJWT(token)
	if err != nil{
		ctx.JSON(getStatusCode(domain.ErrUnathorizedToken), ResponseError{Message: domain.ErrUnathorizedToken.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ResponseProfile{
		Message: domain.SuccessGetProfile,
		User: user,
	})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError, domain.ErrCantSignToken:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnathorizedToken:
		return http.StatusUnauthorized
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}