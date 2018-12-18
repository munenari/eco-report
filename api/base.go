package api

import "net/http"

var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true, // for connection close header
		},
	}
)
