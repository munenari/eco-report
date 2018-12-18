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

// GetBatteryData returns battery json data
func GetBatteryData(origin string, o *model.OneTimePassword, deviceID string) (*map[string]interface{}, error) {
	r := bytes.NewReader([]byte(`{"loadCache":false}`))
	u, err := url.Parse(origin)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "housecontrol/v1/devices", deviceID)
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
	res := new(map[string]interface{})
	decorder := json.NewDecoder(resp.Body)
	err = decorder.Decode(res)
	return res, err
}
