package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	fmt.Println(string(decodeByte))

	// URL encode-decode
	// var data = "https://kalipare.com/"

	// var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	// fmt.Println(encodedString)

	// var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	// var decodedString = string(decodedByte)
	// fmt.Println(decodedString)
}
