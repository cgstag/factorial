package calculus

import "math/big"

func Factorial(n *big.Int) *big.Int{
	x := big.NewInt(1)
	if n.Cmp(big.NewInt(0)) == 0 {
		return x
	}
	return x.Mul(n, Factorial(x.Sub(n, x)))
}
