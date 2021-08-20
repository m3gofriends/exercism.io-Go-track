// Package atbash implements the Atbash cipher.
package atbash

// Atbash implements the Atbash cipher.
func Atbash(s string) string {
	sByte := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			sByte = append(sByte, 'z'-s[i]+'A')

		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			sByte = append(sByte, 'z'-s[i]+'a')

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sByte = append(sByte, s[i])

		default:
			continue
		}

		if len(sByte)%6 == 5 {
			sByte = append(sByte, ' ')
		}
	}

	return string(trim(sByte))
}

func trim(b []byte) []byte {
	if b[len(b)-1] == ' ' {
		b = b[:len(b)-1]
	}

	return b
}
