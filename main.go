package main

import (
	"fmt"
	"math"
	"time"

	"homecredit.vn/prime-go/sieve"
)

var prime_count int64
var SV sieve.Sieve

var mark_i int64
var mark_j int64
var count_i int64
var finished bool

func searchPrime(N int64) {

	fmt.Println("D Initializing sieve with size ", N)
	SV.Init(N)

	SV.Begin()

	nsqrt := int64(math.Sqrt(float64(N)))

	fmt.Println("D Start marking  sqrt(N)=", nsqrt)

	mark_j = 0

	for mark_i = 3; mark_i <= nsqrt; mark_i += 2 {
		byte_i := SV.Get()

		if byte_i != 0 {

			var _2i int64 = 2 * mark_i
			for mark_j = mark_i * mark_i; mark_j <= N; mark_j += _2i {
				SV.Mark(mark_j)
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

	fmt.Println("D Marking done. Counting primes...")

	prime_count = 1
	SV.Begin()
	for count_i = 3; count_i <= N; count_i += 2 {
		byte_i := SV.Get()
		if byte_i != 0 {
			prime_count++
		}
		SV.Next()
	}

	finished = true
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

	finished = false

	go searchPrime(100 * B)

	for !finished {
		fmt.Println("mark: ", mark_i, ">", mark_j, " count: ", count_i)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Found ", prime_count, " primes in ", time.Since(start))
}
