package burningdice

import (
	"math/big"
)

type pascalsTriangle struct {
	rows [][]*big.Int
}

var (
	thePascalsTriangle pascalsTriangle
)

func (pt *pascalsTriangle) addRow() {
	if pt.rows == nil {
		pt.rows = make([][]*big.Int, 1)
		pt.rows[0] = make([]*big.Int, 1)
		pt.rows[0][0] = big.NewInt(1)
		return
	}
	lastRowIdx := len(pt.rows) - 1
	currentLen := len(pt.rows[lastRowIdx]) + 1
	pt.rows = append(pt.rows, make([]*big.Int, currentLen))
	currentRowIdx := lastRowIdx + 1
	
	pt.rows[currentRowIdx][0] = big.NewInt(1);
	for i := 1; (i+1) < currentLen; i++ {
		pt.rows[currentRowIdx][i] = big.NewInt(0).Add(pt.rows[lastRowIdx][i-1], pt.rows[lastRowIdx][i])
	}
	pt.rows[currentRowIdx][currentLen-1] = big.NewInt(1);
}

func (pt *pascalsTriangle) combinations(subTotal, total uint) *big.Int {
	if subTotal > total {
		return big.NewInt(0)
	}
	
	for uint(len(pt.rows)) <= total {
		pt.addRow()
	}
	return pt.rows[total][subTotal]
}

// This free-standing combinations function uses
// a global Pascal's Triangle to cache previous
// combinations answers
func combinations(subTotal, total uint) *big.Int {
	return thePascalsTriangle.combinations(subTotal, total)
}
