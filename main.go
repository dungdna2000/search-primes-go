package main

import (
	"fmt"
	"math"
	"time"

	"homecredit.vn/prime-go/sieve"
)

var prime_count int64

func searchPrime(N int64) {

	sieve.Init(N)
	sieve.Begin()

	nsqrt := int64(math.Sqrt(float64(N)))

	for i := int64(3); i <= nsqrt; i += 2 {

		byte_i := sieve.Get()

		//
		// Found a real prime! process to marking
		//
		if byte_i != 0 {
			for j := i * i; j <= N; j += i {
				if j%2 != 0 {
					sieve.Mark(j)
				}
			}
		}

		sieve.Next()
	}

	fmt.Println("Done marking. Now counting primes!")
	prime_count = 1

	sieve.Begin()
	for i := int64(3); i <= N; i += 2 {
		byte_i := sieve.Get()
		if byte_i != 0 {
			prime_count++ //primes.add(i);
			//fmt.Print(i, " ")
		}

		sieve.Next()
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
	start := time.Now()
	searchPrime(1000000000)
	//searchPrime(100000000000)
	fmt.Println("Found ", prime_count, " primes in ", time.Since(start))
}
