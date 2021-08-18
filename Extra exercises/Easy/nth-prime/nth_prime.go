// Package prime determines what the nth prime number is.
package prime

var primes []int = generatePrimes(104744)

// Nth determines what the nth prime number is.
func Nth(nth int) (int, bool) {
	if nth < 1 {
		return 0, false
	}
	return primes[nth-1], true
}

func generatePrimes(n int) []int {
	numberMap, primeList := make(map[int]bool), make([]int, 0)

	for i := 2; i*i < n; i++ {
		if !numberMap[i] {
			for j := i * i; j < n; j = j + i {
				numberMap[j] = true
			}
		}
	}

	for i := 2; i < n; i++ {
		if !numberMap[i] {
			primeList = append(primeList, i)
		}
	}

	return primeList
}
