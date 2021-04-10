package access_token

import (
	"strings"

	"github.com/Pawelek242/home_oauth-api/src/utils/errors"
)

var error []string

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequest(append(error, "invalid access token"))
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
