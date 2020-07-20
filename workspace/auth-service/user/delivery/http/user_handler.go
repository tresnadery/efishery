package http

import(			
	"time"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"auth-service/domain"
	authHelper "auth-service/helpers/auth"
	validatorHelper "auth-service/helpers/validator"
)

type ResponseError struct{
	Errors interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
}

type ResponseRegister struct{
	Message string `json:"message"`
	Password string `json:"password"`
}
type ResponseToken struct{
	Message string `json:"message"`
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
}
type ResponseProfile struct{
	Message string `json:"message"`
	User interface{} `json:"user"`
}
type UserHandler struct{
	Conn *sql.DB
	uUsecase domain.UserUsecase
}

func NewUserHandler(ctx *gin.Engine, us domain.UserUsecase, conn *sql.DB){
	handler := &UserHandler{
		uUsecase: us,
		Conn: conn,
	}
	ctx.POST("users/registrations", handler.Register)
	ctx.POST("tokens", handler.Token)
	ctx.GET("users/profiles", handler.Profile)
}

func (u *UserHandler) Register(ctx *gin.Context){
	var(
		pass string
		payload domain.User	
	) 
	if err := ctx.BindJSON(&payload); err != nil{
		logrus.Error(err)
		ctx.JSON(getStatusCode(domain.ErrBadParamInput), ResponseError{Message: domain.ErrBadParamInput.Error()})
		return
	}	
	// get user by name and check id user is exists
	user, err := u.uUsecase.GetByNameANDPhoneNumber(payload.Name, payload.PhoneNumber)
	switch{
	case err != nil && getStatusCode(err) == http.StatusInternalServerError:
		logrus.Error(err)
		ctx.JSON(getStatusCode(err), ResponseError{Message:err.Error()})
		return	
	case getStatusCode(err) == http.StatusNotFound:
		validator := validatorHelper.NewValidator(u.Conn)
		if err := validator.IsValidRequest(payload); err != nil{
			logrus.Error(err)
			ctx.JSON(getStatusCode(domain.ErrBadParamInput), ResponseError{
				Message: domain.ErrBadParamInput.Error(),
				Errors: err,
			})
			return	
		}
		pass, err = u.uUsecase.Store(ctx, &payload)
		if err != nil{
			logrus.Error(err)
			ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
	case err == nil:
		pass, err = u.uUsecase.UpdatePassword(ctx, user.ID.String())
		if err != nil{
			logrus.Error(err)
			ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			return
		}
	}	
	ctx.JSON(http.StatusCreated, ResponseRegister{
		Message  : domain.SuccessCreateUser,
		Password : pass, 
	})
	return
}

func (u *UserHandler) Token(ctx *gin.Context){
	var payload domain.PayloadGetToken
	if err := ctx.BindJSON(&payload); err != nil{
		logrus.Error(err)
		ctx.JSON(getStatusCode(domain.ErrBadParamInput), ResponseError{Message: domain.ErrBadParamInput.Error()})
		return
	}
	validator := validatorHelper.NewValidator(u.Conn)
	if err := validator.IsValidRequest(payload); err != nil{
		logrus.Error(err)
		ctx.JSON(getStatusCode(domain.ErrBadParamInput), ResponseError{
			Message: domain.ErrBadParamInput.Error(),
			Errors: err,
		})
		return	
	}
	user, err := u.uUsecase.GetByPhoneNumber(payload.PhoneNumber)
	if err != nil{
		logrus.Error(err)
		ctx.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	
	checkPass :=  bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))		
	if (checkPass != nil && checkPass == bcrypt.ErrMismatchedHashAndPassword){
		logrus.Error(err)
		ctx.JSON(getStatusCode(err), ResponseError{Message: domain.ErrPasswordIsIncorrect.Error()})		
		return 
	}	
	token, err := generateJWTToken(user)
	if err != nil{
		return 
	}
	ctx.JSON(http.StatusOK, ResponseToken{
		Message: domain.SuccessCrateAccessToken,
		AccessToken: token,
		TokenType: "Bearer",
	})
	return
}
func (u *UserHandler) Profile(ctx *gin.Context){
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

func generateJWTToken(payload *domain.User)(string, error){
	expiresAt := time.Now().Add(time.Hour * 1000).Unix()
	tk := &domain.JwtToken{
		PhoneNumber: payload.PhoneNumber,
		Name:  payload.Name,
		RoleName: payload.RoleName,
		CreatedAt: payload.CreatedAt,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	// generate jwt auth
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {		
		return "", domain.ErrCantSignToken
	}	
	return tokenString, nil
}