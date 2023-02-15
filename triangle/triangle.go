package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a > b + c || b > a + c || c > a + b {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
