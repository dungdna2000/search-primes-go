package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"homecredit.vn/prime-go/sieve"
)

// var prime_count int64
//var SV sieve.Sieve

var prime int64
var mark_wait_group sync.WaitGroup
var mark_thread_finished chan int64
var mark_thread_all_finished sync.WaitGroup

var finished bool

func markCompositeOf(sv *sieve.Sieve, N int64, p int64, safe_threshold int64) {
	//fmt.Println(prime, " started!")

	var _2i int64 = 2 * p

	// for composite = p * p; composite <= N; composite += _2i {
	// 	SV.Mark(composite)
	// }

	var cp int64
	for cp = p * p; cp <= safe_threshold; cp += _2i {
		sv.Mark(cp)
	}

	// release lock so that the outer thread can continue
	mark_wait_group.Done()

	// continue to run without worrying about collision!
	for ; cp <= N; cp += _2i {
		sv.Mark(cp)
	}

	mark_thread_all_finished.Done()
	mark_thread_finished <- p
}

func searchPrime(N int64, sv *sieve.Sieve, max_threads int) {

	sv.Init(N)
	sv.Begin()

	nsqrt := int64(math.Sqrt(float64(N)))

	num_threads := 0

	mark_thread_finished = make(chan int64)

	// IDEA: instead of marking composite of different primes, split marking of the same prime?

	var d int64 = 4
	for prime = 5; prime <= nsqrt; prime += d {
		if sv.Get() != 0 {
			// IDEA : if one routine reaches pass safe threshold (?) , we can safely start another thread until max threads reach
			mark_wait_group.Add(1)
			mark_thread_all_finished.Add(1)

			// wait here if enough threads or continue
			if num_threads >= max_threads {
				<-mark_thread_finished // this should block
				num_threads--
			}

			num_threads++
			go markCompositeOf(sv, N, prime, nsqrt)

			mark_wait_group.Wait()
		}

		sv.Next()
		d = flip24(d)
	}

	mark_thread_all_finished.Wait()
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
11	100,000,000,000     	4,118,054,813		~10m
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

func test_case(N int64, expected int64) {
	var sv sieve.Sieve

	searchPrime(N, &sv, 1)

	fmt.Print("Test N=", N, " ... ")
	var prime_count int64 = sv.Count()
	if prime_count == expected {
		fmt.Println("OK")
	} else {
		fmt.Println("FAILED! Actual = ", prime_count, ", Expected = ", expected)

		// re-run with 1 thread to compare
		var sv_expected sieve.Sieve
		searchPrime(N, &sv_expected, 1)

		sv_expected.Compare(&sv)
	}

	//dump_primes(N)
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
	//test_case(10*B, 455052511)

}

func main() {
	start := time.Now()

	finished = false
	//unit_test()

	//test_case(10*B, 455052511)

	//test_case(1000, 168)
	//test_case(10000, 1229)
	//test_case(100000, 9592) // sometimes 1/10 : FAILED! Actual =  9593 , Expected =  9592
	test_case(1000000, 78498)

	//test_case(100000000, 5761455)
	//test_case(B, 50847534)

	// for !finished {
	// 	fmt.Println(time.Since(start), " p: ", prime, ">", composite)
	// 	time.Sleep(2 * time.Second)
	// }

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
