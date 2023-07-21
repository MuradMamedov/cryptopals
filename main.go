package main

import (
	"fmt"
	"strconv"
	"strings"
)

const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	toBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
}

func toBase64(input string) string {
	binaryString := ""
	for _, char := range input {
		h, _ := strconv.ParseInt(string(char), 16, 8)
		b := fmt.Sprintf("%04b", h)
		binaryString = binaryString + b
	}
	fmt.Printf("Binary string: %v, length: %d\n", binaryString, len(binaryString))

	// split to decimal blocks
	var blocks []int64
	padding := 0
	for i := 0; i < len(binaryString); i += 6 {
		rightBound := i + 6
		if len(binaryString) < rightBound {
			padding = (rightBound - len(binaryString)) / 2
			rightBound = len(binaryString)
		}
		b, _ := strconv.ParseInt(binaryString[i:rightBound]+strings.Repeat("0", padding*2), 2, 64)
		blocks = append(blocks, b)
	}
	fmt.Print("Blocks: ")
	for _, number := range blocks {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\nPadding: %d\n", padding)

	// convert to base64
	var output = ""
	for _, number := range blocks {
		output = output + string(base64Chars[number])
	}

	// add padding
	output += strings.Repeat("=", padding)
	fmt.Printf("Output: %v\n", output)
	return output
}
