package sieve

import "fmt"

var bits []byte

func init() {

}

// Internal iterator
var current_byte int
var current_bit int

func Begin() {
	current_byte = 0
	current_bit = 0
}

func Next() {
	current_bit++
	if current_bit == 8 {
		current_byte++
		current_bit = 0
	}
}

// Initialize the sieve to store enough N int64
func Init(N int64) {

	fmt.Println("DEBUG sieve.Init N = ", N)

	required_size := N/16 + 1
	fmt.Println("DEBUG required_size =", required_size)

	bits = make([]byte, required_size)

	len_bits := len(bits)

	fmt.Println("DEBUG len(bits) =", len_bits)

	//
	// initially, set all bits to 1 to tell that they are primes
	//
	for i := 0; i < len_bits; i++ {
		bits[i] = 0b11111111
	}
}

// Retrieve current element.
func Get() byte {
	var v byte
	switch current_bit {
	case 0:
		v = (byte)(bits[current_byte] & 0b00000001)
	case 1:
		v = (byte)(bits[current_byte] & 0b00000010)
	case 2:
		v = (byte)(bits[current_byte] & 0b00000100)
	case 3:
		v = (byte)(bits[current_byte] & 0b00001000)
	case 4:
		v = (byte)(bits[current_byte] & 0b00010000)
	case 5:
		v = (byte)(bits[current_byte] & 0b00100000)
	case 6:
		v = (byte)(bits[current_byte] & 0b01000000)
	case 7:
		v = (byte)(bits[current_byte] & 0b10000000)
	}
	return v
}

// Mark position idx as non prime
func Mark(n int64) {

	// b:	0                    1                      2
	// idx:	0 1 2 3  4  5  6  7  8  9 10 11 12 13 14 15 16 .... (idx)
	// n:	3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 .... (idx*2+3) = ii  => ii = ( idx - 3 ) / 2

	var ii int64 = (n - 3) / 2

	b := int(ii / 8)
	bi := int(ii % 8)

	switch bi {
	case 0:
		bits[b] &= 0b11111110
	case 1:
		bits[b] &= 0b11111101
	case 2:
		bits[b] &= 0b11111011
	case 3:
		bits[b] &= 0b11110111
	case 4:
		bits[b] &= 0b11101111
	case 5:
		bits[b] &= 0b11011111
	case 6:
		bits[b] &= 0b10111111
	case 7:
		bits[b] &= 0b01111111
	}
}
