package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"path"

	"github.com/munenari/eco-report/model"
)

// GetNonce from an api response
func GetNonce(origin string) (*model.Nonce, error) {
	u, err := url.Parse(origin)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "housecontrol/v1/indoorauth/nonce")
	resp, err := client.Get(u.String())
	if err != nil && err != io.EOF {
		fmt.Println("http error:", err)
	}
	defer resp.Body.Close()
	decorder := json.NewDecoder(resp.Body)
	n := new(model.Nonce)
	err = decorder.Decode(n)
	if err != nil {
		return nil, err
	}
	return n, nil
}
