// Package diffiehellman implements Diffie-Hellman key exchange.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a new random private key.
func PrivateKey(p *big.Int) *big.Int {
	a, _ := rand.Int(rand.Reader, big.NewInt(0).Sub(p, big.NewInt(2)))
	return big.NewInt(0).Add(a, big.NewInt(2))
}

// PublicKey generates a new public key.
func PublicKey(private, p *big.Int, g int64) (A *big.Int) {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

// NewPair generates a new private/public key pair.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = big.NewInt(0).Exp(big.NewInt(g), private, p)
	return private, public
}

// SecretKey calculates secret key from own private and strange public key.
func SecretKey(private1, public2, p *big.Int) (s *big.Int) {
	return big.NewInt(0).Exp(public2, private1, p)
}
