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
	primeMap, primeList := make(map[int]bool), []int{2}

	for i := 3; i < n; i += 2 {
		if !primeMap[i] {
			primeList = append(primeList, i)
			for j := i * i; j < n; j += i {
				primeMap[j] = true
			}
		}
	}

	return primeList
}
