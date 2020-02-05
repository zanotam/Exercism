//Package to dtermine what type of triangle something is and test that capability is correct.
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind = int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ Kind = iota // equilateral
	Iso Kind = iota // isosceles
	Sca Kind = iota // scalene
)

// The actual function to determine what type of triangle something is based upon side lengths.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	//NaT if a+b<c for any combination, NaT if a<=0 for any side,
	case (a + b) < c:
		k = NaT
	case (a + c) < b:
		k = NaT
	case (b + c) < a:
		k = NaT
	case (a <= 0) || (b <= 0) || (c <= 0):
		k = NaT
	case !checkInfinite(a, b, c):
		k = NaT
	//Equilateral if a=b=c
	case (a == b) && (b == c):
		k = Equ
	case (a == b) || (b == c) || (a == c):
		k = Iso
	//scalene if a!=b, b!=c, and a!=c
	case (a != b) && (b != c) && (a != c) && !(math.IsNaN(a)) && !(math.IsNaN(b)) && !(math.IsNaN(c)):
		k = Sca
	default:
		k = NaT
	}
	return k
}

func checkInfinite(a, b, c float64) bool {
	return (a != math.Inf(1)) && (b != math.Inf(1)) && (a != math.Inf(-1)) && (b != math.Inf(-1)) && (c != math.Inf(1)) && (c != math.Inf(1))
}
