package main

import (
	"bufio"
	"bytes"
	"encoding/base32"
	"fmt"
	"os"
)

func main() {
	key := []byte{10, 17, 9, 66, 32, 1, 99, 111, 4, 67, 53}

	encryptedText := make([]byte, 0, 10)
	input := bufio.NewScanner(os.Stdin)
	clearText := make([]string, 0)
	output := make([]byte, 4)
	buffer := bytes.NewBuffer(output)
	encoder := base32.NewEncoder(base32.StdEncoding, buffer)

	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}
		clearText = append(clearText, line)
	}

	for _, line := range clearText {
		fmt.Printf("%d %s\n", len(line), line)
		fmt.Printf("%v\n", []byte(line))
		for i := 0; i < len(line); i++ {
			encryptedText = append(encryptedText, line[i]^key[i])
		}
		fmt.Println(encryptedText)
		encoder.Write(encryptedText)
		encryptedText = encryptedText[0:0]
	}
	encoder.Close()
	fmt.Println(buffer)
	fmt.Println()
	fmt.Println("---end---")
}
