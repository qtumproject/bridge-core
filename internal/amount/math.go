package amount

import "math/big"

func Cmp(a, b Amount) int {
	return a.Int().Cmp(b.Int())
}

func Add(a, b Amount) Amount {
	return NewFromBigInt(big.NewInt(0).Add(a.Int(), b.Int()))
}

func Sub(a, b Amount) Amount {
	return NewFromBigInt(big.NewInt(0).Sub(a.Int(), b.Int()))
}

func Mul(a, b Amount) Amount {
	mul := big.NewInt(0).Mul(a.Int(), b.Int())
	return NewFromBigInt(big.NewInt(0).Div(mul, One))
}

func Div(a, b Amount) Amount {
	div := big.NewInt(0).Mul(a.Int(), One)
	return NewFromBigInt(big.NewInt(0).Div(div, b.Int()))
}

func Pow(a Amount, b int64) Amount {
	power := big.NewInt(b)
	numerator := big.NewInt(0).Exp(a.Int(), power, nil)
	denominator := big.NewInt(0).Exp(One, big.NewInt(0).Sub(power, big.NewInt(1)), nil)
	return NewFromBigInt(big.NewInt(0).Div(numerator, denominator))
}

func Sum(amounts ...Amount) Amount {
	sum := NewFromInt(0)
	for _, amount := range amounts {
		sum = Add(sum, amount)
	}
	return sum
}
