package authentication

import (
	"github.com/vongphachan/funny-store-backend/src/modules/admins"
)

type LoginResponse struct {
	User  *admins.Admin `json:"user"`
	Token *string       `json:"token"`
}
