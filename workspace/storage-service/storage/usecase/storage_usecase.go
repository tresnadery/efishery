package usecase

import(
	"storage-service/domain"
)

type StorageUsecase struct{
	StorageRepo domain.StorageRepository
}

func NewStorageUsecase(StorageRepo domain.StorageRepository)*StorageUsecase{
	return &StorageUsecase{StorageRepo}
}

func (r *StorageUsecase) Store(rate float64)(error){	
    return r.StorageRepo.Store(rate)
}

func (r *StorageUsecase) GetRate()(float64, error){
	rate, err := r.StorageRepo.GetRate()
    return rate, err

}