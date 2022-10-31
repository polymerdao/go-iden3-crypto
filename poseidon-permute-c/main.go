package main

import "C"
import (
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// converts three uint64 into one big int and run poseidon
//
//export permute
func permute(elements []uint64) []uint64 {
	arr := (*[12]uint64)(elements)
	res, _ := poseidon.Permute(*arr)
	return res[:]
}

func main() {}
