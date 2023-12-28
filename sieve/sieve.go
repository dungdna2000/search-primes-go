package sieve

import "fmt"

type Sieve struct {
	bits []byte

	current_byte int
	current_bit  int
}

func init() {

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

func (sv Sieve) String() string {
	var res string
	res = fmt.Sprintf("len: %d\n", len(sv.bits))
	for _, b := range sv.bits {
		res += fmt.Sprintf("%08b ", b)
	}
	return res
}

// Initialize the sieve to store enough N int64
func (sv *Sieve) Init(N int64) {

	fmt.Println("DEBUG sieve.Init N = ", N)

	required_size := N/16 + 1
	fmt.Println("DEBUG required_size =", required_size)

	sv.bits = make([]byte, required_size)

	len_bits := len(sv.bits)

	fmt.Println("DEBUG len(bits) =", len_bits)

	//
	// initially, set all bits to 1 to tell that they are primes
	//
	for i := 0; i < len_bits; i++ {
		sv.bits[i] = 0b11111111
	}
}

// Retrieve current element.
func (sv *Sieve) Get() byte {
	var v byte
	switch sv.current_bit {
	case 0:
		v = (byte)(sv.bits[sv.current_byte] & 0b00000001)
	case 1:
		v = (byte)(sv.bits[sv.current_byte] & 0b00000010)
	case 2:
		v = (byte)(sv.bits[sv.current_byte] & 0b00000100)
	case 3:
		v = (byte)(sv.bits[sv.current_byte] & 0b00001000)
	case 4:
		v = (byte)(sv.bits[sv.current_byte] & 0b00010000)
	case 5:
		v = (byte)(sv.bits[sv.current_byte] & 0b00100000)
	case 6:
		v = (byte)(sv.bits[sv.current_byte] & 0b01000000)
	case 7:
		v = (byte)(sv.bits[sv.current_byte] & 0b10000000)
	}
	return v
}

// Mark position idx as non prime
func (sv *Sieve) Mark(n int64) {

	// b:	0                    1                      2
	// idx:	0 1 2 3  4  5  6  7  8  9 10 11 12 13 14 15 16 .... (idx)
	// n:	3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 .... (idx*2+3) = ii  => ii = ( idx - 3 ) / 2

	var ii int64 = (n - 3) / 2

	b := int(ii / 8)
	bi := int(ii % 8)

	switch bi {
	case 0:
		sv.bits[b] &= 0b11111110
	case 1:
		sv.bits[b] &= 0b11111101
	case 2:
		sv.bits[b] &= 0b11111011
	case 3:
		sv.bits[b] &= 0b11110111
	case 4:
		sv.bits[b] &= 0b11101111
	case 5:
		sv.bits[b] &= 0b11011111
	case 6:
		sv.bits[b] &= 0b10111111
	case 7:
		sv.bits[b] &= 0b01111111
	}
}
