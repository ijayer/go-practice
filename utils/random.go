package utils

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"time"
)

// Random return length bit random num
func Random(length int) string {
	var num int
	data := make([]byte, length)

	for i := 0; i < length; i++ {
		num = mrand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = mrand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}

// RandNumMath return 6 bit random num by math
func RandNumMath() string {
	rnd := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	num := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return num
}

// RandNumCrypto return 6 bit random num by crypto
func RandNumCrypto() string {
	//rand.Read
	k := make([]byte, 32)
	if _, err := crand.Read(k); err != nil {
		fmt.Printf("rand.Read() erro : %v \n", err)
	}
	//fmt.Printf("rand.Read(): %v \n", k)

	//rand.Int
	rnd, err := crand.Int(crand.Reader, big.NewInt(1000000))
	if err != nil {
		fmt.Printf("rand.Int() error : %v \n", err)
	}
	num := fmt.Sprintf("%06v", rnd)
	return num
}

// GenerateRandNum 生成最大范围内随机数
func GenerateRandNum() int {
	mrand.Seed(time.Now().Unix())
	randNum := mrand.Intn(100)
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

// GenerateRangeNum 生成一个区间范围的随机数
func GenerateRangeNum(min, max int) int {
	mrand.Seed(time.Now().Unix())
	randNum := mrand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}
