package jwt

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestEasyToken_GetToken(t *testing.T) {
	token, _ := EasyToken{
		Username: strconv.Itoa(100),
		Expires:  time.Now().Unix() + 3600*24*30, //Segundos
	}.GetToken()
	fmt.Println(token)

	b, s, e := EasyToken{}.ValidateToken(token)
	fmt.Println(b,s,e)


}

