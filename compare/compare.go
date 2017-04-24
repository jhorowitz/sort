package compare

import (
	"bytes"
)

type Lexicographical [][]byte

func (p Lexicographical) Len() int           { return len(p) }
func (p Lexicographical) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Lexicographical) Less(i, j int) bool { return bytes.Compare(p[i], p[j]) < 0 }

type Numeric [][]byte

func compareNumeric(a, b []byte) int {
	const zero = byte('0')
	const nine = byte('9')

	const equal = 0
	const aGreater = 1
	const bGreater = -1

	var i int
	for i = 0; i < len(a) && i < len(b); i++ {
		var aIsNum = a[i] <= nine && a[i] >= zero
		var bIsNum = b[i] <= nine && b[i] >= zero
		if aIsNum && !bIsNum {
			return aGreater
		}

		if !bIsNum || !aIsNum {
			if a[i] > b[i] {
				return aGreater
			}
			if a[i] < b[i] {
				return bGreater
			}
			continue
		}
		if bIsNum && !aIsNum {
			return bGreater
		}
		if aIsNum {
			return aGreater
		}
	}

	if len(a) <= i && len(b) <= i {
		return equal
	}

	if len(a) <= i {
		return bGreater
	}

	if len(b) <= i {
		return aGreater
	}
}

func isInt(b byte) bool {


	return false
}

func (p Numeric) Len() int           { return len(p) }
func (p Numeric) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Numeric) Less(i, j int) bool {

	return bytes.Compare(p[i], p[j]) < 0
}
