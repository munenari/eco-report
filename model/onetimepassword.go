package model

import (
	"encoding/base64"
	"fmt"
)

// OneTimePasswordRequest struct
type OneTimePasswordRequest struct {
	Digest string `json:"digest"`
}

// OneTimePassword response struct
type OneTimePassword struct {
	OneTimePassword string `json:"onetimepassword"`
}

// GetAuthHeader by onetime response data
func (o *OneTimePassword) GetAuthHeader() string {
	v := fmt.Sprintf("user:%s", o.OneTimePassword)
	token := base64.StdEncoding.EncodeToString([]byte(v))
	return fmt.Sprintf("Basic %s", token)
}
