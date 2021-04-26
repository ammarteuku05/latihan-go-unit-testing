package main

import (
	"fmt"
	"testing"
)

func TestHitungPajak(t *testing.T) {
	ppn, sum := HitungPajak(110000)
	fmt.Printf("Pajak: %v, Harga: %v\n", ppn, sum)
}
