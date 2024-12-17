package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sureshchandak1/go-orderbook-backend/internal/controllers/auth"
)

func RegisterRoutes() http.Handler {

	r := gin.Default()

	r.POST("/v1/api/auth/signup", auth.Signup)
	r.POST("/v1/api/auth/login", auth.Login)

	return r
}
