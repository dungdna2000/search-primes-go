package main

import (
	"fmt"
	"math"
	"time"

	"homecredit.vn/prime-go/sieve"
)

var prime_count int64
var SV sieve.Sieve

func searchPrime(N int64) {

	fmt.Println("DEBUG Initializing sieve with size ", N)
	SV.Init(N)

	SV.Begin()

	nsqrt := int64(math.Sqrt(float64(N)))

	fmt.Println("DEBUG Start marking  sqrt(N)=", nsqrt)

	var i int64

	for i = 3; i <= nsqrt; i += 2 {

		byte_i := SV.Get()

		if byte_i != 0 {

			var _2i int64 = 2 * i
			for j := i * i; j <= N; j += _2i {
				SV.Mark(j)
			}

			// Use 2 routines to hopefully speed things up!
			// isquare := i * i
			// d := (N - isquare) / 2
			// m := isquare + d
			// m = m - m%i // m%i is to make sure m is at the correct (i-th step)

			// wait_mark.Add(2)
			// go mark(1, isquare, m, i)
			// go mark(2, m+i, N, i)

			// wait_mark.Wait()
		}

		SV.Next()
	}

	fmt.Println("Done marking. Now counting primes!")
	prime_count = 1

	SV.Begin()
	for i = 3; i <= N; i += 2 {
		byte_i := SV.Get()
		if byte_i != 0 {
			prime_count++
		}
		SV.Next()
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
100,000,000,000     4,118,054,813		(ok ~30m)
1,000,000,000,000   37,607,912,018		(   ~300m? )
10,000,000,000,000  346,065,536,839
...
2,000,000,000       98,222,287
*/

const B int64 = 1000000000

func main() {
	start := time.Now()
	//searchPrime(1000000000)

	searchPrime(10000)
	fmt.Println(SV)
	fmt.Println("Found ", prime_count, " primes in ", time.Since(start))
}
