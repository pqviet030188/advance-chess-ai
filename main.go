package main

import (
	"fmt"

	"github.com/pqviet030188/advance-chess-ai/uint96"
	"rsc.io/quote"
)

func main() {
	v1 := uint96.FromUInt64(0xff00ff00ffffffff)
	v2 := uint96.FromUInt64(0xfff0ff00ffffffff)
	zero := uint96.FromUInt32(1)
	fmt.Printf("%32b\n", v1.Lo)
	fmt.Printf("%32b\n", v1.Mid)
	fmt.Printf("%32b\n", v1.Hi)
	fmt.Printf("%t\n", v1.Equals(v2))
	fmt.Printf("%t\n", zero.IsZero())

	big := v2.Big()
	fmt.Printf("0x%96b\n", big)
	// by := byte(0xff)
	// by++

	fmt.Println(quote.Go())
}
