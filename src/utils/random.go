package utils

import (
	"math/rand"
	"time"
	"fmt"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandom(min, max int) int {
	return r.Intn(max-min) + min + 1
}

func GetRandomMax(max int) int {
	return GetRandom(0, max) -1
}

//生成六位的随机数
func GetSixRandom() string {
	return fmt.Sprintf("%06v", r.Int31n(1000000))
}