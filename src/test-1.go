package main

import (
	ReadFile "./read-file"
	"fmt"
	Decoder "./decoder"
)

func main() {
	messages, headers := ReadFile.TextFile("../public/input-file/test.txt")
	decodedMessages := Decoder.Run(headers, messages)

	for _, decodedMessage := range decodedMessages {
		fmt.Println(decodedMessage)
	}
}