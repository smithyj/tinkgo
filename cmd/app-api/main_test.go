package main

import (
	"net/http"
	"testing"
)

func BenchmarkGetAccount(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		_, err := http.Get("http://localhost:8080/account/1")
		if err != nil {
			b.Skipf("%v\n", err.Error())
		}
	}
}