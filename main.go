package main

import (
	"fmt"
	"math"
	"time"

	"sync"

	"homecredit.vn/prime-go/sieve"
)

var prime_count int64

var wait_mark sync.WaitGroup

func mark(id int, start int64, end int64, i int64) {

	// last := time.Now()

	for j := start; j <= end; j += i {
		if j%2 != 0 {
			sieve.Mark(j)
		}

		// if time.Since(last) > time.Second {
		// 	fmt.Println("mark ", id, ": [", i, "]", j, " ")
		// 	last = time.Now()
		// }
	}

	wait_mark.Done()
	fmt.Println("mark ", id, " for ", i, " done!")
}

func searchPrime(N int64) {

	fmt.Println("DEBUG Initializing sieve with size ", N)
	sieve.Init(N)

	sieve.Begin()

	nsqrt := int64(math.Sqrt(float64(N)))

	fmt.Println("DEBUG Marking size Sqrt(N)=", nsqrt)

	var i int64

	//last := time.Now()
	for i = 3; i <= nsqrt; i += 2 {

		byte_i := sieve.Get()

		//
		// Found a real prime! process to marking
		//
		if byte_i != 0 {

			/*
				for j := i * i; j <= N; j += i {
					if j%2 != 0 {
						sieve.Mark(j)
					}

					if time.Since(last) > time.Second {
						fmt.Println("[", i, "]", j, " ")
						last = time.Now()
					}
				}
			*/

			// Use 2 routines to hopefully speed things up!
			//d := (N - i*i) / 2
			//wait_mark.Add(2)
			go mark(1, i*i, N, i)
			//go mark(2, d+i, N, i)

			wait_mark.Wait()

		}

		sieve.Next()

		/*if time.Since(last) > time.Second {
			fmt.Println(i, "<<<<<<<<<<<<<<")
			last = time.Now()
		}*/
	}

	fmt.Println("Done marking. Now counting primes!")
	prime_count = 1

	sieve.Begin()
	for i = 3; i <= N; i += 2 {
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
10,000,000,000      455,052,511         (ok)      459,176,864   10,000,000,000
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

	searchPrime(10 * B)
	fmt.Println("Found ", prime_count, " primes in ", time.Since(start))
}
