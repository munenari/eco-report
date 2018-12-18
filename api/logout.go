package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/munenari/eco-report/model"
)

// Logout request will send
func Logout(origin string, o *model.OneTimePassword) error {
	u, err := url.Parse(origin)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "housecontrol/v1/indoorauth/logout")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", o.GetAuthHeader())
	resp, err := client.Do(req)
	if err != nil && err != io.EOF {
		fmt.Println("http error:", err)
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	fmt.Println("logouted, status:", resp.StatusCode)
	return err
}
