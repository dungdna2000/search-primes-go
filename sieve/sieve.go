package sieve

import (
	"fmt"
)

type Sieve struct {
	bits []byte

	current_byte int
	current_bit  int

	N int64
}

var mask_get [8]byte
var mask_mark [8]byte

func init() {
	mask_get = [8]byte{
		0b00000001,
		0b00000010,
		0b00000100,
		0b00001000,
		0b00010000,
		0b00100000,
		0b01000000,
		0b10000000,
	}

	mask_mark = [8]byte{
		0b11111110,
		0b11111101,
		0b11111011,
		0b11110111,
		0b11101111,
		0b11011111,
		0b10111111,
		0b01111111,
	}

}

func (sv *Sieve) Begin() {
	sv.current_byte = 0
	sv.current_bit = 0
}

func (sv *Sieve) Next() {
	sv.current_bit++
	if sv.current_bit == 8 {
		sv.current_byte++
		sv.current_bit = 0
	}
}

// Initialize the sieve to store enough N int64
func (sv *Sieve) Init(N int64) {
	sv.N = N
	required_size := N/(3*8) + 1
	sv.bits = make([]byte, required_size)
	len_bits := len(sv.bits)

	//
	// initially, set all bits to 1 to tell that they are primes
	//
	for i := 0; i < len_bits; i++ {
		sv.bits[i] = 0b11111111
	}
}

// Retrieve current element.
func (sv *Sieve) Get() byte {
	return sv.bits[sv.current_byte] & mask_get[sv.current_bit]

	// var v byte
	// switch sv.current_bit {
	// case 0:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b00000001)
	// case 1:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b00000010)
	// case 2:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b00000100)
	// case 3:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b00001000)
	// case 4:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b00010000)
	// case 5:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b00100000)
	// case 6:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b01000000)
	// case 7:
	// 	v = (byte)(sv.bits[sv.current_byte] & 0b10000000)
	// }
	// return v
}

// Mark position idx as non prime
func (sv *Sieve) Mark(n int64) {

	//
	// compact sieve for 2
	//
	//var ii int64 = (n - 3) / 2

	//
	// compact sieve for 2, 3
	//
	var ii int64 = (n - 5) / 3
	m := (n - 5) % 3
	if m == 2 {
		ii++
	} else if m != 0 {
		return // m==1 is a multiple of 3 , already remove > skip !
	}

	b := int(ii / 8)
	bi := int(ii % 8)

	if b > len(sv.bits) {
		fmt.Println(">> panic! n = ", n, "ii = ", ii, " ", b)
	}

	sv.bits[b] = sv.bits[b] & mask_mark[bi]

	// switch bi {
	// case 0:
	// 	sv.bits[b] &= 0b11111110
	// case 1:
	// 	sv.bits[b] &= 0b11111101
	// case 2:
	// 	sv.bits[b] &= 0b11111011
	// case 3:
	// 	sv.bits[b] &= 0b11110111
	// case 4:
	// 	sv.bits[b] &= 0b11101111
	// case 5:
	// 	sv.bits[b] &= 0b11011111
	// case 6:
	// 	sv.bits[b] &= 0b10111111
	// case 7:
	// 	sv.bits[b] &= 0b01111111
	// }
}

func flip24(d int64) int64 {
	if d == 4 {
		return 2
	}
	return 4
}

func (sv *Sieve) Count() int64 {
	var prime int64 = 0
	var count int64 = 2
	var d int64 = 4
	sv.Begin()
	for prime = 5; prime <= sv.N; prime += d {
		if sv.Get() != 0 {
			count++
		}
		sv.Next()
		d = flip24(d)
	}
	return count
}

func (sv *Sieve) Compare(svOther *Sieve) {
	var prime int64
	var d int64 = 4

	sv.Begin()
	svOther.Begin()

	for prime = 5; prime <= sv.N; prime += d {
		if (sv.Get() != 0) != (svOther.Get() != 0) {
			fmt.Println(prime)
		}
		sv.Next()
		svOther.Next()
		d = flip24(d)
	}
}
