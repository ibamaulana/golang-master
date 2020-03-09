package services

import (
	"github.com/ibamaulana/golang-master/model"
	"github.com/jinzhu/gorm"
)

type PenilaianServiceContract interface {
	Get() ([]*model.Penilaian, error)
}

type penilaianContractService struct {
	db *gorm.DB
}

func NewPenilaianServiceContract(db *gorm.DB) PenilaianServiceContract {
	return &penilaianContractService{db}
}

func (srv *penilaianContractService) Get() ([]*model.Penilaian, error) {
	var user []*model.Penilaian
	var err error

	err = srv.db.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
