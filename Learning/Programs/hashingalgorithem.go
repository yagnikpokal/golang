package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func main() {
	fmt.Println(GetMD5Hash("hello"))                                        // Fist and second requirement imposible to extrac has function and hash message
	fmt.Println(GetMD5Hash("cello"))                                        // Third requirement Small change in message will produce drastic differenece
	fmt.Println(GetMD5Hash("The quick brown fox jumped over the lazy dog")) // Fourth requirement digest/ouput have same length
	//Wether you take one string, 10 string, full book string the length will be same

}
