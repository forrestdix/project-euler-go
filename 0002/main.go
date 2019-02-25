package main

import "log"

func main() {
	evens := evenFibonacci(4*1000*1000)
	sum := int64(0)
	for i := 0; i < len(evens); i++ {
		sum += evens[i]
	}
	log.Printf("sum == %d", sum)
}

func evenFibonacci(max int64) (results []int64) {
	odd1 := int64(1)
	even := int64(2)
	odd2 := int64(3)

	for even <= max {
		results = append(results, even)
		odd1 = even + odd2
		even = odd2 + odd1
		odd2 = odd1 + even
	}

	return
}