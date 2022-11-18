package main

import (
	"encoding/base32"
	"fmt"

	//"os"
	"bytes"
)

func main() {
	clearText := []byte("Kuł")
	key := []byte{10, 17, 9, 66}
	encryptedText := make([]byte, 4)
	fmt.Println(string(clearText))
	fmt.Println(clearText)
	fmt.Println(key)
	for i := 0; i < len(clearText); i++ {
		encryptedText[i] = clearText[i] ^ key[i]
	}
	fmt.Println(encryptedText)
	output := make([]byte, 4)
	buffer := bytes.NewBuffer(output)
	encoder := base32.NewEncoder(base32.StdEncoding, buffer)
	encoder.Write(encryptedText)
	encoder.Close()
	for i := 0; i < len(encryptedText); i++ {
		clearText[i] = encryptedText[i] ^ key[i]
	}
	fmt.Println(clearText)
	fmt.Println(string(clearText))
	fmt.Printf("%v\n", []byte("aaa"))
	fmt.Println(buffer)
}
