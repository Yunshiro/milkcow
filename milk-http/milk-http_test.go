package milk_http

import (
	"testing"
)

func TestGET(t *testing.T) {
	GET("/get", ":9001", Show())
}

func Show() MilkFunc {
	dst := "hello"
	return MilkTransfer(dst)
}
