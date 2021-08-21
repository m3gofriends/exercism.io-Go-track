// Package cipher implements a simple shift cipher like Caesar and a more secure substitution cipher.
package cipher

var (
	upperCase = [26]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	lowerCase = [26]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

type (
	newShiftStruct struct {
		distance int
	}

	newVigenereStruct struct {
		key string
	}
)

// NewCaesar returns a Caesar cipher, which is a shift cipher of distance 3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a shift cipher of a given distance.
func NewShift(distance int) Cipher {
	if (distance < 1 || distance > 25) && (distance < -25 || distance > -1) {
		return nil
	}
	return newShiftStruct{distance}
}

// NewVigenere returns a Vigenère cipher with a given key.
func NewVigenere(key string) Cipher {
	counter := 0

	for i := 0; i < len(key); i++ {
		if key[i] == 'a' {
			counter++
		} else if key[i] < 'a' || key[i] > 'z' {
			return nil
		}
	}

	if counter == len(key) {
		return nil
	}

	return newVigenereStruct{key}
}

func (n newShiftStruct) Encode(s string) string {
	sByte := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			sByte = append(sByte, lowerCase[(searchIndex(s[i], upperCase)+(26+n.distance))%26])
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			sByte = append(sByte, lowerCase[(searchIndex(s[i], lowerCase)+(26+n.distance))%26])
		}
	}

	return string(sByte)
}

func (n newShiftStruct) Decode(s string) string {
	sByte := []byte(s)

	for i := 0; i < len(sByte); i++ {
		sByte[i] = lowerCase[(searchIndex(sByte[i], lowerCase)+(26-n.distance))%26]
	}

	return string(sByte)
}

func (n newVigenereStruct) Encode(s string) string {
	s, n.key = reshape(s, n.key)
	sByte := []byte(s)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			sByte[i] = lowerCase[(searchIndex(s[i], upperCase)+searchIndex(n.key[i], lowerCase))%26]
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			sByte[i] = lowerCase[(searchIndex(s[i], lowerCase)+searchIndex(n.key[i], lowerCase))%26]
		}
	}

	return string(sByte)
}

func (n newVigenereStruct) Decode(s string) string {
	s, n.key = reshape(s, n.key)
	sByte := []byte(s)

	for i := 0; i < len(sByte); i++ {
		sByte[i] = lowerCase[(searchIndex(sByte[i], lowerCase)+(26-searchIndex(n.key[i], lowerCase)))%26]
	}

	return string(sByte)
}

func reshape(s, key string) (string, string) {
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') {
			s = s[:i] + s[i+1:]
		}
	}

	for i := len(key) - 1; i >= 0; i-- {
		if (key[i] < 'A' || key[i] > 'Z') && (key[i] < 'a' || key[i] > 'z') {
			key = key[:i] + key[i+1:]
		}
	}

	if len(s) > len(key) {
		keyByte := []byte(key)
		for i := 0; len(s) > len(keyByte); i++ {
			keyByte = append(keyByte, key[i%len(key)])
		}
		key = string(keyByte)
	}

	return s, key
}

func searchIndex(b byte, array [26]byte) (index int) {
	for {
		if b == array[index] {
			return index
		}
		index++
	}
}
