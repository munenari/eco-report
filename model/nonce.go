package model

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// Nonce response struct
type Nonce struct {
	Nonce          string `json:"nonce"`
	PasswordExists bool   `json:"passwordExists"`
}

// GetDigest by nonce
func (n *Nonce) GetDigest(passwd string) string {
	md5Data := md5.Sum([]byte(fmt.Sprintf("%s:%s", n.Nonce, passwd)))
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", md5Data[:])))
}
