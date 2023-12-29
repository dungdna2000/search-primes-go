package main

import (
	"fmt"
	"math"
	"time"

	"homecredit.vn/prime-go/sieve"
)

var prime_count int64
var SV sieve.Sieve

var p int64         // the current found prime
var composite int64 // the composite to be marked
var count_i int64
var finished bool

func searchPrime(N int64) {

	SV.Init(N)
	SV.Begin()

	nsqrt := int64(math.Sqrt(float64(N)))

	//fmt.Println("D Start marking  sqrt(N)=", nsqrt)

	composite = 0
	prime_count = 2

	var d int64 = 4
	for p = 5; p <= nsqrt; p += d {
		if SV.Get() != 0 {
			var _2i int64 = 2 * p
			for composite = p * p; composite <= N; composite += _2i {
				SV.Mark(composite)
			}
			//prime_count++
		}

		SV.Next()
		d = flip24(d)
	}

	prime_count = 2
	d = 4
	SV.Begin()
	for count_i = 5; count_i <= N; count_i += d {
		if SV.Get() != 0 {
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
09	1,000,000,000       	50,847,534			~7s
10	10,000,000,000      	455,052,511         ~1m24s
11	100,000,000,000     	4,118,054,813		~0m
12	1,000,000,000,000   	37,607,912,018		<< this is our TARGET
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

func test_case(N int64, expected int64) {
	searchPrime(N)
	fmt.Print("Test N=", N, " ... ")
	if prime_count == expected {
		fmt.Println("OK")
	} else {
		fmt.Println("FAILED! Actual = ", prime_count, ", Expected = ", expected)
	}
}

func unit_test() {
	test_case(100, 25)
	test_case(1000, 168)
	test_case(10000, 1229)
	test_case(100000, 9592)
	test_case(1000000, 78498)
	test_case(10000000, 664579)
	test_case(100000000, 5761455)
	test_case(B, 50847534)
}

func main() {
	start := time.Now()

	finished = false
	unit_test()
	//go test_case(10*B, 455052511)
	//go test_case(1000*B, 4118054813)

	for !finished {
		fmt.Println(time.Since(start), " p: ", p, ">", composite, " count: ", count_i)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Total time:", time.Since(start))
}

// func main1() {
// 	var N int64 = 100
// 	var d int64 = 4
// 	var p int64

// 	for p = 5; p < N; p += d {
// 		fmt.Printf("%02d ", p)
// 		d = flip24(d)
// 	}
// 	fmt.Println()

// 	d = 4
// 	for p = 5; p < N; p += d {

// 		t := (p - 5) / 3
// 		if (p-5)%3 == 2 {
// 			t += 1
// 		}

// 		fmt.Printf("%02d ", t)
// 		d = flip24(d)
// 	}
// 	fmt.Println()

// }
