package util

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//随机数种子
var Rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func VerifySign(body, key string) string {
	return strings.ToLower(MD5(strings.ToLower(MD5(body)) + key))
}

//md5加密
func MD5(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])
}

func MD5Encrypt16(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])[8:24]
}

func GetAbsPath(relativePath string) (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]) + relativePath)
}

// 获取数字随机字符
func GetRandDigit(n int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(n)+"d", Rnd.Intn(int(math.Pow10(n))))
}

// 创建命令 数据
func CreateCmd(cmd byte, context []byte,uid int32) []byte {
	msg := make([]byte, 0)
	// 包头
	msg = append(msg, 0x7f)
	//len(context) + 6
	// 命令号
	msg = append(msg, cmd)

	// 长度  2个字节
	toBytes, _ := IntToBytes(int16(len(context) + 10))
	msg = append(msg, toBytes...)

	// uid  4个字节
	uidBytes, _ := IntToBytes(uid)
	msg = append(msg, uidBytes...)

	// 内容
	msg = append(msg, context...)
	//校验和
	var check int
	for _, v := range msg {
		check += int(v)
	}
	check = check + 0x8f
	msg = append(msg, byte(check%(math.MaxUint8+1)))

	//包尾
	msg = append(msg, 0x8f)
	return msg
}

//整形转换成字节  大端模式   高位在前
func IntToBytes(n interface{}) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes(), err
}

//字节转换成整形
func BytesToInt(n interface{},b []byte) error {
	bytesBuffer := bytes.NewBuffer(b)
	err := binary.Read(bytesBuffer, binary.BigEndian, n)
	return err
}


//校验包合法性 true 不合法
func IsCheckCmd(cmd []byte,uid int32) bool {

	len := len(cmd)
	//最小包长度
	if len < 10 {
		return false
	}
	//包头包尾 7f 8f
	if cmd[0] != 0x7f || cmd[len-1] != 0x8f {
		return false
	}
	// uid
	var cmdUid int32
	err := BytesToInt(&cmdUid,cmd[4:10])
	if uid != cmdUid || err != nil {
		return false
	}

	//包长度
	var cmdLen int16
	err = BytesToInt(&cmdLen,cmd[2:4])
	if len != int(cmdLen) || err != nil {
		return false
	}



	//校验和
	var check int
	for i, v := range cmd {
		if i != len-2 {
			check += int(v)
		}
	}

	if byte(check%(math.MaxUint8+1)) != cmd[len-2] {
		return false
	}
	return true
}
