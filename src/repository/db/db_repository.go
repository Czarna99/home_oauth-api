package db

import (
	"github.com/Pawelek242/home_oauth-api/src/domain/access_token"
	"github.com/Pawelek242/home_oauth-api/src/utils/errors"
)

var error []string

func NewRepository() DbRepository {
	return &dbRepository{}

}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError(append(error, "database connection not implemented yet"))

}
