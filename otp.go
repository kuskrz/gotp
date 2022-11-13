package main

import (
	"bufio"
	// "encoding/base32"
	"fmt"
	"os"
)

func main() {
	//key := []byte{10, 17, 9, 66, 32, 1, 99, 111, 4, 67, 53}
	// encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	// encryptedText := make([]byte, 0, 10)
	input := bufio.NewScanner(os.Stdin)
	clearText := make([]string, 0)

	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}
		clearText = append(clearText, line)
	}

	for _, line := range clearText {
		fmt.Printf("%v\n", []byte(line))
	}
	/* fmt.Printf("%d %s\n", len(line), line)
	for i := 0; i < len(line); i++ {
		encryptedText = append(encryptedText, line[i]^key[i])
	}
	fmt.Println(encryptedText)
	encoder.Write(encryptedText)
	fmt.Println()
	encryptedText = encryptedText[0:0] */
	// encoder.Close()
	fmt.Println("qq")
}
