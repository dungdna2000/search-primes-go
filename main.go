package main

import (
	"fmt"
	"math"
)

var prime_count int64

func searchPrime(N int64) {

	prime_bits := make([]byte, N/16+1)

	//
	// initially, set all odd numbers to primes
	//
	for i := 0; i < len(prime_bits); i++ {
		prime_bits[i] = 0b11111111
	}

	nsqrt := int64(math.Sqrt(float64(N)))

	// current bit and byte position so that we can "move" along with i each step
	pos_byte := int(0)
	pos_bit := int(0)

	for i := int64(3); i < nsqrt; i++ {
		byte_i := byte(0b11111111)
		switch pos_bit {
		case 0:
			byte_i = (byte)(prime_bits[pos_byte] & 0b00000001)
		case 1:
			byte_i = (byte)(prime_bits[pos_byte] & 0b00000010)
		case 2:
			byte_i = (byte)(prime_bits[pos_byte] & 0b00000100)
		case 3:
			byte_i = (byte)(prime_bits[pos_byte] & 0b00001000)
		case 4:
			byte_i = (byte)(prime_bits[pos_byte] & 0b00010000)
		case 5:
			byte_i = (byte)(prime_bits[pos_byte] & 0b00100000)
		case 6:
			byte_i = (byte)(prime_bits[pos_byte] & 0b01000000)
		case 7:
			byte_i = (byte)(prime_bits[pos_byte] & 0b10000000)
		}

		//
		// Found a real prime! process to marking
		//
		if byte_i != 0 {
			for j := i * i; j <= N; j += i {
				if j%2 != 0 {
					jj := (j - 3) / 2

					b := int(jj / 8)
					bi := int(jj % 8)

					switch bi {
					case 0:
						prime_bits[b] &= 0b11111110
					case 1:
						prime_bits[b] &= 0b11111101
					case 2:
						prime_bits[b] &= 0b11111011
					case 3:
						prime_bits[b] &= 0b11110111
					case 4:
						prime_bits[b] &= 0b11101111
					case 5:
						prime_bits[b] &= 0b11011111
					case 6:
						prime_bits[b] &= 0b10111111
					case 7:
						prime_bits[b] &= 0b01111111
					}
				}
			}
		}

		pos_bit++
		if pos_bit == 8 {
			pos_bit = 0
			pos_byte++
		}
	}

	fmt.Println("Done marking. Now counting primes!")
	prime_count = 1
	pos_byte = 0
	pos_bit = 0
	byte_i := byte(0b11111111)
	for i := int64(3); i <= N; i += 2 {
		// extract bit at position corresponding to interger i
		switch pos_bit {
		case 0:
			byte_i = prime_bits[pos_byte] & 0b00000001
		case 1:
			byte_i = prime_bits[pos_byte] & 0b00000010
		case 2:
			byte_i = prime_bits[pos_byte] & 0b00000100
		case 3:
			byte_i = prime_bits[pos_byte] & 0b00001000
		case 4:
			byte_i = prime_bits[pos_byte] & 0b00010000
		case 5:
			byte_i = prime_bits[pos_byte] & 0b00100000
		case 6:
			byte_i = prime_bits[pos_byte] & 0b01000000
		case 7:
			byte_i = prime_bits[pos_byte] & 0b10000000
		}

		if byte_i != 0 {
			prime_count++ //primes.add(i);
		}

		pos_bit++
		if pos_bit == 8 {
			pos_bit = 0
			pos_byte++
		}
	}

}

/*

https://t5k.org/howmany.html

100			        25
1000		        168
10000		        1,229
100,000		        9,592
1,000,000	        78,498
10,000,000          664,579
100,000,000         5,761,455
1,000,000,000       50,847,534
10,000,000,000      455,052,511         (ok)
100,000,000,000     4,118,054,813
1,000,000,000,000   37,607,912,018
10,000,000,000,000  346,065,536,839
...
2,000,000,000       98,222,287
*/

func main() {
	searchPrime(10000)
	fmt.Println(prime_count)
}
