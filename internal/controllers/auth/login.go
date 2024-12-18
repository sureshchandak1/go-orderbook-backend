package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sureshchandak1/go-orderbook-backend/internal/database"
	"github.com/sureshchandak1/go-orderbook-backend/internal/database/tables"
	"github.com/sureshchandak1/go-orderbook-backend/internal/dtos"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var body dtos.LoginReqDto
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

	// Look up requested user
	var user tables.User
	database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECERET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.GeneralError(err))
		return
	}

	// send response
	data := dtos.LoginResData{Token: tokenString}
	response := dtos.LoginResDto{
		BaseResponse: *dtos.GetBaseResponse(http.StatusOK, "login successfully"),
		LoginResData: data,
	}

	c.JSON(http.StatusOK, response)

}
