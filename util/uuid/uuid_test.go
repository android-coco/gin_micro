package uuid

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	uuid := NewUUID()
	hex := uuid.Hex()
	hex32 := uuid.Hex32()
	upper := uuid.HexToUpper()
	fmt.Print(hex,"\n",hex32,"\n",upper)
}
