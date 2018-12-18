package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/munenari/eco-report/model"
)

// GetInstantData returns instantdata buf
func GetInstantData(origin string, o *model.OneTimePassword, filterValue string) (*model.InstantData, error) {
	r := bytes.NewReader([]byte(filterValue))
	u, err := url.Parse(origin)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "housecontrol/v1/log/hems/instantdata")
	req, err := http.NewRequest("POST", u.String(), r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", o.GetAuthHeader())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil && err != io.EOF {
		fmt.Println("http error:", err)
	}
	defer resp.Body.Close()
	i := new(model.InstantData)
	decorder := json.NewDecoder(resp.Body)
	err = decorder.Decode(i)
	if err != nil {
		return nil, err
	}
	return i, nil
}
