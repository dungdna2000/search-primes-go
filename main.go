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

	var d int64 = 4
	for mark_i = 5; mark_i <= nsqrt; mark_i += d {
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
		d = flip24(d)
	}

	fmt.Println("D Marking done. Counting primes...")

	prime_count = 2
	SV.Begin()
	for count_i = 5; count_i <= N; count_i += d {
		byte_i := SV.Get()
		if byte_i != 0 {
			prime_count++
		}
		SV.Next()
		d = flip24(d)
	}

	finished = true
}

/*

https://t5k.org/howmany.html

01 	10						4
02	100			        	25
03	1000		        	168
04	10000		        	1,229
05	100,000		        	9,592
06	1,000,000	        	78,498
07	10,000,000          	664,579
08	100,000,000         	5,761,455
09	1,000,000,000       	50,847,534			9s
10	10,000,000,000      	455,052,511         1m42s
11	100,000,000,000     	4,118,054,813
12	1,000,000,000,000   	37,607,912,018		(300m)
13	10,000,000,000,000  	346,065,536,839
14	100,000,000,000,000		3,204,941,750,802
15	1,000,000,000,000,000	29,844,570,422,669

*/

const B int64 = 1000000000

func flip24(d int64) int64 {
	if d == 4 {
		return 2
	}
	return 4
}

// func main() {
// 	N := 100

// 	var d int
// 	d = 4
// 	for i := 5; i < N; i += d {
// 		fmt.Printf("%02d ", i)
// 		d = flipflop(d, 2, 4)
// 	}
// 	fmt.Println()

// 	d = 4
// 	for i := 5; i < N; i += d {

// 		t := (i - 5) / 3
// 		if (i-5)%3 == 2 {
// 			t += 1
// 		}

// 		fmt.Printf("%02d ", t)
// 		d = flipflop(d, 2, 4)
// 	}
// 	fmt.Println()

// }

func main() {
	start := time.Now()

	finished = false

	searchPrime(1000)

	// for !finished {
	// 	fmt.Println(time.Since(start), ": P: ", mark_i, ">", mark_j, " count: ", count_i)
	// 	time.Sleep(2 * time.Second)
	// }

	fmt.Println("Found ", prime_count, " primes in ", time.Since(start))
}
