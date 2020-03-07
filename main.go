package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ibamaulana/golang-master/controller/auth"
	"github.com/ibamaulana/golang-master/controller/ping"
	"github.com/ibamaulana/golang-master/controller/users"
	"github.com/ibamaulana/golang-master/jwtmiddleware"
)

func main() {
	r := gin.Default()

	signInKey := "secret"
	jwtmiddleware.InitJWTMiddlewareCustom([]byte(signInKey), jwt.SigningMethodHS512)

	r.Use(jwtmiddleware.CORSMiddleware())
	r.GET("ping", ping.PingController)

	{
		userRoute := r.Group("user")
		userRoute.Use(jwtmiddleware.MyAuth())

		userRoute.POST("", users.CreateController)
		userRoute.GET(":id", users.FindController)
		userRoute.GET("", users.FindByController)
	}

	{
		authRoute := r.Group("auth")
		authRoute.POST("", auth.LoginController)
	}

	r.Run(":9000")
}
