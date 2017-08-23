package main
import "crypto/sha512"
//import "fmt"

func Sha512Hash(input string) []byte {
	sha_512 := sha512.New()
	sha_512.Write([]byte(input))
	return sha_512.Sum(nil)
}