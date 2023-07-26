package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	for i := 0; i < 256; i++ {
		decoded := decypherXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", fmt.Sprintf("%x", i))
		bs, _ := hex.DecodeString(decoded)
		fmt.Printf("%x %v\n", i, string(bs)) // Cooking MC's like a pound of bacon, key - X (x58)
	}
	// xor := xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	// fmt.Printf("Xor: %v\n", xor)
	// output := toBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	// fmt.Printf("Output: %v\n", output)
}

func decypherXor(input string, char string) string {
	return xor(input, strings.Repeat(char, len(input)))
}

func xor(input1 string, input2 string) string {
	// X xor Y -> (X || Y) && !(X && Y)
	output := ""
	for i := 0; i < len(input1); i++ {
		h1, _ := strconv.ParseInt(string(input1[i]), 16, 8)
		h2, _ := strconv.ParseInt(string(input2[i]), 16, 8)
		xor := (h1 | h2) & ^(h1 & h2)
		output += strconv.FormatInt(xor, 16)
	}

	return output
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
	return output
}
