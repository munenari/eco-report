package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"path"

	"github.com/munenari/eco-report/model"
)

// GetOneTimePassword returns a onetimepassword response struct
func GetOneTimePassword(origin string, n *model.Nonce, password string) (*model.OneTimePassword, error) {
	d := &model.OneTimePasswordRequest{
		Digest: n.GetDigest(password),
	}
	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	u, err := url.Parse(origin)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "housecontrol/v1/indoorauth/onetimepassword")
	resp, err := client.Post(u.String(), "application/x-www-form-urlencoded", r)
	if err != nil && err != io.EOF {
		fmt.Println("http error:", err)
	}
	defer resp.Body.Close()
	decorder := json.NewDecoder(resp.Body)
	p := new(model.OneTimePassword)
	err = decorder.Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
