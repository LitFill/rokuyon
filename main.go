/*
rokuyon LitFill <email>
*/
package main

import (
	"encoding/base64"
	"fmt"
	"log/slog"
	"os"
)

func encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func decode(input string) string {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return string(decoded)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("    encode <input>")
	fmt.Println("    decode <input>")
}

func main() {
	if len(os.Args) < 3 {
		slog.Error("Subcommand diperlukan (encode/decode) dan input string.")
		os.Exit(1)
	}

	subcommand := os.Args[1]
	inputString := os.Args[2]

	switch subcommand {
	case "encode":
		fmt.Println(encode(inputString))
	case "decode":
		fmt.Println(decode(inputString))
	default:
		slog.Error("subcommand invalid, gunakan `encode` atau `decode`")
		os.Exit(1)
	}
}
