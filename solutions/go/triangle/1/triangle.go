package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 || !(a+b >= c) || !(b+c >= a) || !(a+c >= b) {
		return Kind(0)
	} else if a == b && b == c {
		return Kind(1)
	} else if a == b || b == c || c == a {
		return Kind(2)
	} else {
		return Kind(3)
	}
}