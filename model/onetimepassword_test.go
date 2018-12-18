package model

import "testing"

func TestGetAuthHeader(t *testing.T) {
	o := &OneTimePassword{OneTimePassword: "test_password"}
	expected := "Basic dXNlcjp0ZXN0X3Bhc3N3b3Jk"
	authHeader := o.GetAuthHeader()
	if expected != authHeader {
		t.Error("string was not match, expected:", expected, "actual:", authHeader)
	}
}
