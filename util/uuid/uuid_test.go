/*
 * @Author: yhlyl
 * @Date: 2019-10-22 14:13:08
 * @LastEditTime: 2019-11-04 16:30:39
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /gin_micro/util/uuid/uuid_test.go
 */
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
	fmt.Print(hex, "\n", hex32, "\n", upper)
}
