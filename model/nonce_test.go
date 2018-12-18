package model

import "testing"

func TestGetDigest(t *testing.T) {
	n := &Nonce{Nonce: "12345678"}
	password := "my_password"
	expected := "YTJiYmFmZDc4Y2FhMzcxZjc5NjJkOTljMjU0N2M2MzU="
	digest := n.GetDigest(password)
	if expected != digest {
		t.Error("digest was not match, expected:", expected, "actual", digest)
	}
}
