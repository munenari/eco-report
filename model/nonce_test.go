package model

import (
	"bytes"
	"io"
	"io/ioutil"
	"runtime"
	"strings"
	"testing"
)

func TestGetDigest(t *testing.T) {
	n := &Nonce{Nonce: "12345678"}
	password := "my_password"
	expected := "YTJiYmFmZDc4Y2FhMzcxZjc5NjJkOTljMjU0N2M2MzU="
	digest := n.GetDigest(password)
	if expected != digest {
		t.Error("digest was not match, expected:", expected, "actual", digest)
	}
}

func BenchmarkCast(b *testing.B) {
	buf := []byte(strings.Repeat("sampledata", 1024))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(buf)
		io.Copy(ioutil.Discard, r)
		// log.Println(string(buf))
		runtime.GC()
	}
}

// func BenchmarkReader(b *testing.B) {
// 	buf := make([]byte, 1024*1024)
// 	b.Error(cap(buf))
// 	return
// 	b.SetParallelism(8)
// 	b.N = 1000
// 	b.ResetTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			r := bytes.NewReader(buf)
// 			io.Copy(ioutil.Discard, r)
// 			time.Sleep(1 * time.Second)
// 		}
// 	})
// }
