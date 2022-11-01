package main

import "C"
import (
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// converts three uint64 into one big int and run poseidon
//
//export permute
func permute(e0, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11 uint64) (uint64, uint64, uint64, uint64, uint64, uint64,
	uint64, uint64, uint64, uint64, uint64, uint64) {
	res, _ := poseidon.Permute([12]uint64{e0, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11})
	return res[0], res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10], res[11]
}

func main() {}
