// Package sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
package sieve

// Sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
func Sieve(n int) (primes []int) {
	primeMap := make(map[int]bool)

	for i := 2; i <= n; i++ {
		if !primeMap[i] {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				primeMap[j] = true
			}
		}
	}

	return primes
}
