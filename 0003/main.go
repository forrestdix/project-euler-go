package main

import (
	"log"
	"math"
)

func main() {
	largest := largestPrimeFactor(600851475143, []int64{2})
	log.Printf("largest: %d", largest)
}

func largestPrimeFactor(x int64, primes []int64) (factor int64) {
	limit := int64(math.Sqrt(float64(x)))

	for i := 0; i < len(primes); i++ {
		if x%(x/primes[i]) == 0 {
			return largestPrimeFactor(x/primes[i], primes)
		}
	}

	for primes[len(primes) - 1] <= limit {
		primes = append(primes, nextPrime(primes[len(primes) - 1]))
		if x%(x/primes[len(primes) - 1]) == 0 {
			return largestPrimeFactor(x/primes[len(primes) - 1], primes)
		}
	}

	return x
}

func nextPrime(after int64) int64 {
	for {
		after++
		if isPrime(after) {
			break
		}
	}

	return after
}

func isPrime(value int64) bool {
	for i := int64(2); i <= int64(math.Sqrt(float64(value))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}