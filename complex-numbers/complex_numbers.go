package complexnumbers

import "math"

type Number struct {
	r, i float64
}

func (n Number) Real() float64 {
	return n.r
}

func (n Number) Imaginary() float64 {
	return n.i
}

func (n1 Number) Add(n2 Number) Number {
	var res Number
	res.r = n1.r + n2.r
	res.i = n1.i + n2.i
	return res
}

func (n1 Number) Subtract(n2 Number) Number {
	var res Number
	res.r = n1.r - n2.r
	res.i = n1.i - n2.i
	return res
}

func (n1 Number) Multiply(n2 Number) Number {
	var res Number
	res.r = n1.r * n2.r - n1.i * n2.i
	res.i = n1.r * n2.i + n1.i * n2.r
	return res
}

func (n Number) Times(factor float64) Number {
	var res Number
	res.r = n.r * factor
	res.i = n.i * factor
	return res
}

func (n1 Number) Divide(n2 Number) Number {
	var res Number
	var div float64
	div = n2.r * n2.r + n2.i * n2.i
	res.r = (n1.r * n2.r + n1.i * n2.i) / div
	res.i = (n1.i * n2.r - n1.r * n2.i) / div
	return res
}

func (n Number) Conjugate() Number {
	var res Number
	res.r = n.r
	res.i = -n.i
	return res
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.r * n.r + n.i * n.i)
}

func (n Number) Exp() Number {
	var res Number
	exp := math.Exp(n.r)
	res.r = exp * math.Cos(n.i)
	res.i = exp * math.Sin(n.i)
	return res
}
