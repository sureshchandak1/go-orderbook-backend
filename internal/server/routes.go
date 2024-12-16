package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sureshchandak1/go-orderbook-backend/internal/controllers/auth"
)

func RegisterRoutes() http.Handler {

	r := gin.Default()

	r.POST("/signup", auth.Signup)
	r.POST("/login", auth.Login)

	return r
}
