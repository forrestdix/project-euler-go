package main

import "log"

func main() {
	multiplesOf3 := multiplesOf(3, 1000)
	multiplesOf5 := multiplesOf(5, 1000)
	uniqueMultiples := map[int64]bool{}

	for i := 0; i < len(multiplesOf3); i++ {
		uniqueMultiples[multiplesOf3[i]] = true
	}
	for i := 0; i < len(multiplesOf5); i++ {
		uniqueMultiples[multiplesOf5[i]] = true
	}

	sum := int64(0)
	for k := range uniqueMultiples {
		sum += k
	}

	log.Printf("sum == %d", sum)
}

func multiplesOf(x int64, below int64) (multiples []int64) {
	for i := (below - 1) / x; i > 0; i-- {
		multiples = append(multiples, i * x)
	}

	return
}