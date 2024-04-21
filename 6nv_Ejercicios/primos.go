package main

import (
	"fmt"
	"math"
)

/**
* Accepts a range of integers [min, max] and
* returns an array with all the prime numbers in this range.
*
* Execution is done sequentially
 */

func isPrime(candidate int64) bool {
	var i, limit int64

	if candidate == 2 {
		return true
	}

	if math.Mod(float64(candidate), 2) == 0 {
		return false
	}

	limit = int64(math.Ceil(math.Sqrt(float64(candidate))))
	//fmt.Println("limit:", limit)

	for i = 3; i <= limit; i += 2 { //Only odd numbers

		if math.Mod(float64(candidate), float64(i)) == 0 {
			return false
		}
	}
	return true
}

func primesInRange(min int64, max int64) []int64 {
	var primeNumbers []int64

	for i := min; i <= max; i++ {
		if isPrime(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}

func main() {
	//xi := primesInRange(1, 100)
	//fmt.Println(xi)
	xi := primesInRangeParallel(1, 1000000)
	fmt.Println(xi)

}

type PrimeResult struct {
	number int64 // A number
	prime  bool  // Is prime or not
}

func firePrimeCalculations(min int64, max int64, channel chan PrimeResult) {
	var i int64
	for i = min; i <= max; i++ {
		go isPrimeAsync(i, channel)
	}
}

func isPrimeAsync(number int64, channel chan PrimeResult) {

	result := new(PrimeResult)
	result.number = number
	result.prime = isPrime(number)
	channel <- *result
}

func primesInRangeParallel(min int64, max int64) []int64 {
	var primeNumbers []int64
	var res PrimeResult
	//var prev int64

	channel := make(chan PrimeResult)
	defer close(channel)

	go firePrimeCalculations(min, max, channel)

	//prev = 0
	for i := min; i <= max; i++ {
		res = <-channel
		if res.prime {
			primeNumbers = append(primeNumbers, res.number)

			/*done := 100 * (i - min) / (max - min)
			if prev != done {
				fmt.Printf("%d %% done.\n", done)
				prev = done
			}
			*/
		}
	}
	return primeNumbers
}
