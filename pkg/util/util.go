package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RandomNumberic(size int) string {
	if size <= 0 {
		panic(fmt.Sprintf("{ size: %d } must be more than 0", size))
	}
	
	r := rand.New(rand.NewSource(time.Now().Unix()))
	
	value := ""
	for i := 0; i < size; i++ {
		value += strconv.Itoa(r.Intn(10))
	}

	return value
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	if n <= 0 {
		panic(fmt.Sprintf("{ size: %d } must be more than 0", n))
	}
	
	r := rand.New(rand.NewSource(time.Now().Unix()))

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}


// EndOfDay 当天的最后一刻
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}