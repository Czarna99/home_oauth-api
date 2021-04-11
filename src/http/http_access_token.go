package http

import (
	"net/http"

	"github.com/Pawelek242/home_oauth-api/src/domain/access_token"
	"github.com/Pawelek242/home_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

var error []string

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}

}

func (handler *accessTokenHandler) GetById(c *gin.Context) {

	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)

}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequest(append(error, "Invalid JSON body"))
		c.JSON(0, restErr) // fix JSON error
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(0, err) // fix JSON error
		return
	}
	c.JSON(http.StatusCreated, at)
}
