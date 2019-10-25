package royalpay

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes2 := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes2[r.Intn(len(bytes2))])
	}
	return string(result)
}

func sign(partnerCode, credentialCode, nonceStr, strTime string) string {
	strParam := fmt.Sprintf(
		"%s&%s&%s&%s",
		partnerCode,
		strTime,
		nonceStr,
		credentialCode)
	sha := sha256.New()
	sha.Write([]byte(strParam))
	sign := sha.Sum(nil)
	return fmt.Sprintf("%x", sign)
}

func getTime() string {
	now := time.Now()
	strTime := strconv.Itoa(int(now.UnixNano()))
	strTime = strTime[:13]
	return strTime
}
