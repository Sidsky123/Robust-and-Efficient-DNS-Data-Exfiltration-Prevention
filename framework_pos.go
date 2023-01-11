// domain data encoding/decoding algo for FrameworkPOS Malware DNS-Tunneling Variant,
// as described on:
// https://blog.gdata.de/artikel/neue-variante-von-frameworkpos-schoepft-daten-ueber-dns-anfragen-ab/
//

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

const A = 0xAA
const B = 0x9B
const C = 0xC3

var op string
var in string

func init() {
	flag.StringVar(&op, "op", "decode", "operation to perform: 'encode' or 'decode', default is 'decode'")
	flag.StringVar(&in, "in", "", "Inputstring to perform choosen operation on")
}

// decode - decodes single values from dns queries into cleartext
func decode(data []byte) (result []byte) {
	for _, v := range data {
		a := v ^ A
		b := a ^ B
		result = append(result, b^C)
	}
	return result
}

// encode - encodes cleartext values to be used in dns queries
func encode(data []byte) (result []byte) {
	for _, v := range data {
		b := v ^ C
		a := b ^ B
		result = append(result, a^A)
	}
	return result
}

func main() {
	flag.Parse()
	input := []byte(in)
	switch {
	case op == "encode":
		encoded := encode(input)
		fmt.Printf("%s %x\n", input, string(encoded))
	case op == "decode":
		data, err := hex.DecodeString(string(input))
		if err != nil {
			panic(err)
		}
		decoded := decode(data)
		fmt.Printf("%x %s\n", data, string(decoded))
	default:
		prog := os.Args[0]
		fmt.Printf("For USAGE INFO call: '%s -h'\n", prog)
	}
}
