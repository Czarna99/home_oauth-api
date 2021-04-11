package access_token

import (
	"strings"
	"time"

	"github.com/Pawelek242/home_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequest(append(error, "Invalid access token ID"))
	}
	if at.UserID <= 0 {
		return errors.NewBadRequest(append(error, "Invalid user ID"))
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequest(append(error, "Invalid client ID"))
	}
	if at.Expires <= 0 {
		return errors.NewBadRequest(append(error, "Invalid expiration date"))
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {

	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
