package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"

	"homecredit.vn/prime-go/sieve"
)

var prime int64
var nsqrt int64
var mark_wait_group sync.WaitGroup

var finished bool

func markSegment(sv *sieve.Sieve, from int64, to int64, delta int64) {
	for cp := from; cp <= to; cp += delta {
		sv.Mark(cp)
	}

	mark_wait_group.Done()
}

func markCompositeOf(sv *sieve.Sieve, N int64, p int64, num_threads int64) {
	var _2p int64 = 2 * p
	var pp int64 = p * p

	// for composite = p * p; composite <= N; composite += _2i {
	// 	SV.Mark(composite)
	// }

	// we will divide range p*p to N into many segments, each will be marked by one thread
	var d int64 = (N - pp) / num_threads

	// seg_size must be divisble by _2p
	seg_size := (d - d%_2p)

	var t int64
	var from int64 = pp

	for t = 1; t <= num_threads; t++ {

		var to int64 = from + seg_size

		// if p is large enough and seg_size is small enough, this could really happens
		//
		if to > N {
			to = N
			t = num_threads + 1 // no more threads!
		}

		// if last thread does not reach to the last N, let it be!
		if t == num_threads {
			if to < N {
				to = N
			}
		}

		mark_wait_group.Add(1)
		go markSegment(sv, from, to, _2p)

		from = to + _2p
	}
	mark_wait_group.Wait()
}

func searchPrime(N int64, sv *sieve.Sieve, num_threads int64) {

	sv.Init(N)
	sv.Begin()

	nsqrt = int64(math.Sqrt(float64(N)))

	var d int64 = 4
	for prime = 5; prime <= nsqrt; prime += d {
		if sv.Get() != 0 {
			markCompositeOf(sv, N, prime, num_threads)
		}

		sv.Next()
		d = flip24(d)
	}

}

const B int64 = 1000000000

func flip24(d int64) int64 {
	if d == 4 {
		return 2
	}
	return 4
}

// func dump_primes(N int64) {
// 	var d int64 = 4
// 	var p int64
// 	SV.Begin()
// 	for p = 5; p <= N; p += d {
// 		if SV.Get() != 0 {
// 			fmt.Print(p, " ")
// 		}
// 		SV.Next()
// 		d = flip24(d)
// 	}
// 	fmt.Println()
// }

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
10	10,000,000,000      	455,052,511         ~1m24s	~13s ~16.0444363s
11	100,000,000,000     	4,118,054,813		~10m  - 2m5.23s

    200,000,000,000			8,007,105,059		4m30 (!verified)
	300,000,000,000         11,818,439,135		6m59 (!verified)
	500,000,000,000         19,308,136,142      12m7 (!verified)

12	1,000,000,000,000   	37,607,912,018		<< this is our TARGET
13	10,000,000,000,000  	346,065,536,839
14	100,000,000,000,000		3,204,941,750,802
15	1,000,000,000,000,000	29,844,570,422,669

*/

func test_case(N int64, expected int64, num_threads int64) {
	var sv sieve.Sieve

	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Println("| Search primes N=", N, ". Threads = ", num_threads)
	fmt.Println("---------------------------------------------------------------------------------")

	start := time.Now()

	searchPrime(N, &sv, num_threads)

	var log string
	f, _ := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	var prime_count int64 = sv.Count()
	if prime_count == expected {
		log = fmt.Sprintln("\n", N, " PASSED! Actual=Expected=", prime_count, ".Time: ", time.Since(start))
	} else {
		log = fmt.Sprintln("\n", N, "FAILED! Actual=", prime_count, ",Expected=", expected, ".Time: ", time.Since(start))

		// re-run with 1 thread to compare
		//var sv_expected sieve.Sieve
		//searchPrime(N, &sv_expected, 1)

		//sv_expected.Compare(&sv)
	}

	f.WriteString(log)
	fmt.Print(log)

	finished = true

	//dump_primes(N)
}

func all_test_cases(threads int64) {
	test_case(100, 25, threads)
	test_case(1000, 168, threads)
	test_case(10000, 1229, threads)
	test_case(100000, 9592, threads)
	test_case(1000000, 78498, threads)
	test_case(10000000, 664579, threads)
	test_case(100000000, 5761455, threads)
	test_case(B, 50847534, threads)
	//test_case(10*B, 455052511, threads)
}

func main() {
	//start := time.Now()

	finished = false

	all_test_cases(20)

	//test_case(10*B, 455052511)

	//go test_case(1000, 168)
	//go test_case(10000, 1229)
	//go test_case(100000, 9592)
	//go test_case(1000000, 78498)
	//go test_case(100000000, 5761455, 10)
	//go test_case(B, 50847534, 10)
	//go test_case(10*B, 455052511, 10)
	//go test_case(100*B, 4118054813, 10)
	//go test_case(B, 50847534, 10)
	//go test_case(100*B, 4118054813, 20)
	//go test_case(500*B, 4118054813, 20)

	//go test_case(1000*B, 37607912018, 20)

	// lastp := prime

	// for !finished {
	// 	if prime > lastp {
	// 		fmt.Print("\n", time.Since(start), ":", nsqrt, ",", prime)
	// 		lastp = prime
	// 	} else {
	// 		fmt.Print(".")
	// 	}

	// 	time.Sleep(5 * time.Second)
	// }

}

// func main_skip_3() {
// 	var N int64 = 200
// 	var d int64 = 4
// 	var p int64

// 	for p = 5; p <= N; p+=d {
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

// var d_pattern = [8]int64{2, 1, 2, 1, 2, 3, 1, 3}

// p+=d - to generate next number in seq
var d_pattern = [8]int64{4, 2, 4, 2, 4, 6, 2, 6}

// m = (p-7)*4 % 15
//var m_pattern = [8]int64{0, 1, 9, 10, 3, 4, 13, 6}

//var primes []int64

func main_skip35() {
	var N int64 = 4000
	var p int64

	// for p = 0; p < N; p++ {
	// 	fmt.Printf("%4d ", p)
	// }
	// fmt.Println()

	// primes = make([]int64, 0)

	var i int64 = 0
	for p = 7; p <= N; p += 2 {

		if p%3 != 0 && p%5 != 0 {
			fmt.Printf("%3d ", i)
			i++
			// 	fmt.Printf("%3d ", p)
		}
	}
	fmt.Println()

	//d_idx := -1
	//j := int64(0)

	// for p = 7; p <= N; p += d_pattern[d_idx] {
	// 	fmt.Printf("%3d ", (p-7)*4/15)

	// 	d_idx++
	// 	if d_idx == 8 {
	// 		d_idx = 0
	// 	}
	// 	j++
	// }

	// fmt.Println()

	d_idx := -1
	for p = 7; p <= N; p += d_pattern[d_idx] {

		var j int64 = ((p - 7) * 4) / 15
		var m int64 = ((p - 7) * 4) % 15

		// magic seq of m :  0, 1, 9, 10, 3, 4, 13, 6

		if m > 1 {
			j++
		}

		fmt.Printf("%3d ", m)

		d_idx++
		if d_idx == 8 {
			d_idx = 0
		}
		j++
	}

	// fmt.Println("DONE!")

	//fmt.Println()

	// for p = 7; p <= N; p += 2 {
	// 	if p%3 != 0 && p%5 != 0 {
	// 		fmt.Printf("%3d ", p*13/48)
	// 	}
	// }
	// fmt.Println()

	// pp := int64(0)
	// prev := int64(7)
	// i := 0
	// for p = 7; p <= N; p += 2 {

	// 	if p%3 != 0 && p%5 != 0 {

	// 		//j := p * 1 / 4

	// 		fmt.Printf("%4d ", (p-prev)/2)
	// 		prev = p

	// 		if i%8 == 0 {
	// 			fmt.Println()
	// 		}
	// 		i++

	// 	}
	// 	pp++

	// }
	// fmt.Println()
}
