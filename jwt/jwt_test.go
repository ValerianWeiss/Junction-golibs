package jwt

import (
	"os"
	"testing"
)

var (
	sub    = "auth0|1234567890"
	userid = "1234567890"
	name   = "John Doe"
	num    = "25.42"
	token  = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhdXRoMHwxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJ1c2VyaWQiOjE2ODQ5NjUxLCJudW0iOjI1LjQyfQ.-Fl-I0na2mWfUunATTLH52TwNxEsaJqODvtuNvGXWOE"
)

func init() {
	os.Setenv("Http_Authorization", token)
}

func TestGetClaimFloat64(t *testing.T) {
	result, err := GetClaim("num")
	checkError(t, err)
	checkResult(t, result, num)
	t.Log(result)
}

func TestGetClaimString(t *testing.T) {
	result, err := GetClaim("name")
	checkError(t, err)
	checkResult(t, result, name)
	t.Log(result)
}

func TestGetClaimSub(t *testing.T) {
	result, err := GetClaimSub()
	checkError(t, err)
	checkResult(t, result, sub)
	t.Log(result)
}

func TestGetUserid(t *testing.T) {
	result, err := GetUserid()
	checkError(t, err)
	checkResult(t, result, userid)
	t.Log(result)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func checkResult(t *testing.T, result, expected string) {
	if result != expected {
		t.Error("value is invalid")
		t.Fail()
	}
}
