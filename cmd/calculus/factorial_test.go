package calculus

import (
	"math/big"
	"testing"
)

func Test_Factorial(t *testing.T) {
	var n = big.NewInt(int64(5))
	x := new(big.Int)
	x.SetString("120",10)
	if Factorial(n).Cmp(x) != 0 {
		t.Errorf("Factorial method incorrect, got: %s, want: %s.", x.String(), "120")
	}
	n = big.NewInt(int64(40))
	x = new(big.Int)
	x.SetString("815915283247897734345611269596115894272000000000",10)
	if Factorial(n).Cmp(x) != 0 {
		t.Errorf("Factorial method incorrect, got: %s, want: %s.", x.String(), "815915283247897734345611269596115894272000000000")
	}
}