package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sureshchandak1/go-orderbook-backend/internal/database/tables"
	"github.com/sureshchandak1/go-orderbook-backend/internal/dtos"
)

func GetUser(c *gin.Context) {

	userAny, exits := c.Get("user")
	if !exits {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var user tables.User = userAny.(tables.User)
	userDto := dtos.UserDto{
		Id:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		MobileNo: user.MobileNo,
	}

	c.JSON(http.StatusOK, userDto)

}
