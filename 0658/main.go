package main

import "log"

func main() {
	log.Printf("I(3, 0) = %d", I(3,0))
	log.Printf("I(3, 2) = %d", I(3,2))
	log.Printf("I(3, 4) = %d", I(3,4))
	log.Printf("S(4, 4) = %d", S(4,4))
	log.Printf("S(8, 8) = %d", S(8,8))
	log.Printf("S(10, 100) = %d", S(10,100))
}

func S(k int64, l int64) (sum int64) {
	for ki := int64(1); ki <= k; ki++ {
		sum += I(ki, l)
		log.Printf("I(%d, %d) = %d", ki, l, I(ki, l))
	}

	return
}

func I(a int64, l int64) (sum int64) {
	sum = 1
	for li := int64(1); li <= l; li++ {
		sum += pow(a, li) - permutationsWithAllCharacters(a, li)
	}

	return
}

func permutationsWithAllCharacters(a int64, l int64) (result int64) {
	permutations := pow(a, l)
	for ai := a - 1; ai > 0; ai-- {
		permutations -= permutationsWithAllCharacters(ai, l) * (factorial(a) / factorial(a - 1))
	}
	return permutations
}

func pow(x int64, y int64) (result int64) {
	if y == 0 {
		return 1
	} else {
		result = x
	}

	for i := int64(1); i < y; i++ {
		result *= x
	}

	return
}

func factorial(x int64) (result int64) {
	if x < 0 {
		return 0
	}

	result = x

	for x > 1 {
		x--
		result *= x
	}

	return
}