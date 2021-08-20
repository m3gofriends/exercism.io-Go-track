// Package rotationalcipher implements the Caesar cipher.
package rotationalcipher

var (
	upperCase = [26]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	lowerCase = [26]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

// RotationalCipher implements the Caesar cipher.
func RotationalCipher(s string, shift int) string {
	if shift == 0 || shift == 26 {
		return s
	}

	sByte := []byte(s)
	for i := 0; i < len(sByte); i++ {
		switch sByte[i] {
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			sByte[i] = upperCase[(searchIndex(sByte[i], upperCase)+shift)%26]
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			sByte[i] = lowerCase[(searchIndex(sByte[i], lowerCase)+shift)%26]
		}
	}

	return string(sByte)
}

func searchIndex(b byte, array [26]byte) (index int) {
	for {
		if b == array[index] {
			return index
		}
		index++
	}
}
