package domain

import(
	"context"
	"time"	
	"github.com/satori/go.uuid"
)

type User struct{
	ID uuid.UUID `json:"id"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=17,is_phone_number_exists"`
	Name string `json:"name" validate:"required,min=5,max=50"`
	RoleID string `json:"role_id" validate:"required,uuid4"`
	RoleName string `json:"role_name"`
	Password string `json:"-"`
	RememberToken string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`	
}

type PayloadGetToken struct{
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUsecase interface{
	GetByPhoneNumber(phoneNumber string)(*User, error)
	GetByNameANDPhoneNumber(name, phoneNumber string)(*User, error)
	Store(ctx context.Context, payload *User)(string, error)
	UpdatePassword(ctx context.Context, userID string)(string, error)
}

type UserRepository interface{
	GetByPhoneNumber(phoneNumber string)(*User, error)
	GetByNameANDPhoneNumber(name, phoneNumber string)(*User, error)
	Store(ctx context.Context, payload *User)(error)
	UpdatePassword(ctx context.Context, userID string, password string)error
}