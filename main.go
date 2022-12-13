package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: cidr-to-mask CIDR")
	os.Exit(-1)
}

func printInvalid(input string) {
	fmt.Printf("Invalid input '%s'\n", input)
	os.Exit(-1)
}

func main() {
	if len(os.Args) > 2 {
		printUsage()
		return
	}
	input := ""
	if len(os.Args) == 2 {
		input = os.Args[1]
	} else {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		input = strings.TrimSpace(string(data))
	}
	cidr, err := strconv.Atoi(input)
	if err != nil || cidr < 0 || cidr > 32 {
		printInvalid(input)
		return
	}
	mask := uint32(0b11111111_11111111_11111111_11111111)
	mask = (mask << (32 - cidr)) & mask
	octet0 := uint8(mask >> 24)
	octet1 := uint8(mask >> 16)
	octet2 := uint8(mask >> 8)
	octet3 := uint8(mask)
	fmt.Printf("%d.%d.%d.%d\n", octet0, octet1, octet2, octet3)
}
