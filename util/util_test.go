package util

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	//var (
	//	in       = "1"
	//	expected = "c4ca4238a0b923820dcc509a6f75849b"
	//)
	//actual := MD5(in)
	//if actual != expected {
	//	t.Errorf("MD5(%s) = %s; expected %s", in, actual, expected)
	//}

	//var x uint64 = 1
	//bytes , err:= IntToBytes(x)
	//fmt.Print(bytes,err)

	//i, err := BytesToInt16([]byte{0,1})
	//fmt.Println(i,err)

	//var x = []byte{1,2,4,5}
	//fmt.Print(x[3:4])
	var json = `{
    "app_version": "1",
    "app_name": "1",
    "num_register_origin": 1,
    "machine_id": "11",
    "account": "18822855253",
    "mobile": "18822855253",
    "password": "123456",
    "agent_id": 10
}`
	xx := VerifySign(base64.StdEncoding.EncodeToString([]byte(json)), DesKey)
	fmt.Println(xx)
}
