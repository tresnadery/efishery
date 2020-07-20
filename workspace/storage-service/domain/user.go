package domain

import(	
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