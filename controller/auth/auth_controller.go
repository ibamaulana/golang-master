package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibamaulana/golang-master/config"
	"github.com/ibamaulana/golang-master/httpresponse"
	"github.com/ibamaulana/golang-master/jwtmiddleware"
	"github.com/ibamaulana/golang-master/request/auth"
	"github.com/ibamaulana/golang-master/services"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(ctx *gin.Context) {
	var err error
	cfg := config.NewConfig()

	db, err := config.MysqlConnection(cfg)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	var req auth.LoginRequest
	if err = ctx.ShouldBind(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	userContract := services.NewUserServiceContract(db)
	user, err := userContract.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	//	 compare password from db with request
	byteHash := []byte(user.Password) // password from db
	bytePlain := []byte(req.Password) // password from request
	
	if err := bcrypt.CompareHashAndPassword(byteHash, bytePlain); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusForbidden, err)
		return
	}

	tokenStruct := new(jwtmiddleware.TokenRequestStructure)
	tokenStruct.Email = user.Email
	tokenStruct.UserID = user.ID

	signInKey := "secret"
	g := jwtmiddleware.NewCustomAuth([]byte(signInKey))
	token, err := g.GenerateToken(*tokenStruct)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusForbidden, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, token)
	return

}
