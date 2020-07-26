package main

import (
	b64 "encoding/base64"
	. "fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	Println(string(sDec))

	Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	Println(string(uDec))
}
