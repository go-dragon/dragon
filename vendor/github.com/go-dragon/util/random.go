package util

import (
	trueRand "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var numbers = []rune("0123456789")

// 是否播随机种子, 用于解决多进程随机种子相同的情况
var seedFlag = false

// check rand seed
func checkSeed() {
	if !seedFlag {
		// 如果没有播随机种子
		rand.Seed(time.Now().UnixNano())
		seedFlag = true
	}
}

// RandomStr generate random string with pseudo-random,maybe 0 begin
func RandomStr(length int) string {
	checkSeed()
	runes := make([]rune, length)
	lettersLen := len(letters)
	for i := 0; i < length; i++ {
		runes[i] = letters[rand.Intn(lettersLen)]
	}
	return string(runes)
}

// RandomNumber generate random number string, not 0 start with pseudo-random
func RandomNumber(length int) string {
	checkSeed()
	runes := make([]rune, length)
	numbersLen := len(numbers)
	runes[0] = numbers[1+rand.Intn(numbersLen-1)]
	for i := 1; i < length; i++ {
		runes[i] = numbers[rand.Intn(numbersLen)]
	}
	return string(runes)
}

// TrueRandomStr generate random string with true random,maybe 0 begin
func TrueRandomStr(length int) string {
	runes := make([]rune, length)
	lettersLen := len(letters)
	// use crypto/rand,for real random
	for i := 0; i < length; i++ {
		n, _ := trueRand.Int(trueRand.Reader, big.NewInt(int64(lettersLen)))
		runes[i] = letters[int(n.Int64())]
	}
	return string(runes)
}

// TrueRandomNumber generate random number string, not 0 start with true random
func TrueRandomNumber(length int) string {
	runes := make([]rune, length)
	numbersLen := len(numbers)
	runes[0] = numbers[1+rand.Intn(numbersLen-1)]
	// 初始化第一个rune
	for i := 1; i < length; i++ {
		// use crypto/rand,for real random
		n, _ := trueRand.Int(trueRand.Reader, big.NewInt(int64(numbersLen)))
		runes[i] = numbers[int(n.Int64())]
	}
	return string(runes)
}
