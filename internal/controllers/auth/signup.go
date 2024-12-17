package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sureshchandak1/go-orderbook-backend/internal/database"
	"github.com/sureshchandak1/go-orderbook-backend/internal/database/tables"
	"github.com/sureshchandak1/go-orderbook-backend/internal/dtos"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var body dtos.SignUpReqDto
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.GeneralError(err))
		return
	}

	// request validation
	if err := validator.New().Struct(body); err != nil {
		validateErrs := err.(validator.ValidationErrors)
		c.JSON(http.StatusInternalServerError, dtos.ValidationError(validateErrs))
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.SignUpReqData.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.GeneralError(err))
		return
	}

	user := tables.User{Name: body.Name, Email: body.Email, Password: string(hash), MobileNo: body.MobileNo}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dtos.GeneralError(err))
		return
	}

	response := getSignUpResponse(user)

	// Respond
	c.JSON(http.StatusOK, response)
}

func getSignUpResponse(user tables.User) *dtos.SignUpResDto {
	data := &dtos.SignUpResData{
		Id:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		MobileNo: user.MobileNo,
	}
	return &dtos.SignUpResDto{
		BaseResponse:  *dtos.GetBaseResponse(http.StatusOK, "successfully register user"),
		SignUpResData: *data,
	}
}
