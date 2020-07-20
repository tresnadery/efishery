package usecase

import(		
	"time"
	"context"	
	"auth-service/domain"
	"auth-service/helpers/randomize"
	auth "auth-service/helpers/auth"
)

type UserUsecase struct{
	userRepo domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) *UserUsecase{
	return &UserUsecase{
		userRepo: u,
		contextTimeout: timeout,
	}
}
func generatePassword()(string, string, error){
	pass := randomize.RandStringBytes(4)
	encodePass, err := auth.HashAndSaltPassword([]byte(pass))
	if err != nil{
		return "", "", domain.ErrInternalServerError
	}
	return pass, encodePass, nil
}
func (u *UserUsecase) GetByPhoneNumber(phoneNumber string)(*domain.User, error){
	user, err := u.userRepo.GetByPhoneNumber(phoneNumber)
	if err != nil{
		return nil, err
	}
	return user, err
}

func (u *UserUsecase) GetByNameANDPhoneNumber(name, phoneNumber string)(*domain.User, error){
	user, err := u.userRepo.GetByNameANDPhoneNumber(name, phoneNumber)
	if err != nil{
		return nil, err
	}
	return user, err
}

func (u *UserUsecase) Store(ctx context.Context, payload *domain.User)(string, error){
	pass, encodePass, err := generatePassword()
	if err != nil{
		return pass, err
	}
	payload.Password = encodePass
	if err := u.userRepo.Store(ctx, payload); err != nil{
		return "", err
	}
	return pass, nil
}

func (u *UserUsecase) UpdatePassword(ctx context.Context, userID string)(string, error){
	pass, encodePass, err := generatePassword()
	if err != nil{
		return pass, err
	}	
	if err := u.userRepo.UpdatePassword(ctx, userID, encodePass); err != nil{
		return "", err
	}
	return pass, nil
}