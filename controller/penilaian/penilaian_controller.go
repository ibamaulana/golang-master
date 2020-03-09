package penilaian

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ibamaulana/golang-master/config"
	"github.com/ibamaulana/golang-master/httpresponse"
	"github.com/ibamaulana/golang-master/services"
)

func GetController(ctx *gin.Context) {
	runtime.GOMAXPROCS(2)

	var err error
	cfg := config.NewConfig()

	db, err := config.MysqlConnection(cfg)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	penilaianContract := services.NewPenilaianServiceContract(db)
	data, err := penilaianContract.Get()

	httpresponse.NewSuccessResponse(ctx, data)
	return
}
