package users

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ibamaulana/golang-master/config"
	"github.com/ibamaulana/golang-master/httpresponse"
	"github.com/ibamaulana/golang-master/model"
	"github.com/ibamaulana/golang-master/request/users"
	"github.com/ibamaulana/golang-master/services"
	"github.com/jinzhu/copier"
)

func CreateController(ctx *gin.Context) {
	runtime.GOMAXPROCS(1)

	var err error
	cfg := config.NewConfig()

	db, err := config.MysqlConnection(cfg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req users.CrateRequest

	if err = ctx.ShouldBindWith(&req, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userContract := services.NewUserServiceContract(db)

	tx := db.Begin()
	defer tx.Rollback()

	user := new(model.User)
	err = copier.Copy(&user, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	passTemp, err := config.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Password = string(passTemp)

	err = userContract.Create(user, tx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})

	return
}

func FindController(ctx *gin.Context) {
	runtime.GOMAXPROCS(2)

	var err error
	cfg := config.NewConfig()

	db, err := config.MysqlConnection(cfg)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	var req users.FindRequest

	if err = ctx.ShouldBindUri(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	userContract := services.NewUserServiceContract(db)
	data, err := userContract.Find(req.ID)

	httpresponse.NewSuccessResponse(ctx, data)
	return
}

func FindByController(ctx *gin.Context) {
	var err error
	cfg := config.NewConfig()

	db, err := config.MysqlConnection(cfg)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	var req users.FindRequest

	if err = ctx.ShouldBind(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	userContract := services.NewUserServiceContract(db)

	data, err := userContract.FindBy(nil)

	if err = ctx.ShouldBind(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, data)
	return
}
