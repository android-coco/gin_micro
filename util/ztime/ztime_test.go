package ztime

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	minute := EndOfMinute(time.Now())
	fmt.Println(minute)
	pow10 := math.Pow10(6)
	fmt.Println(pow10)
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	i := rnd.Intn(int(pow10))
	fmt.Println(fmt.Sprintf("%d",i))
}
