package burningdice

import (
	"math/big"
)

type exponentCache struct {
	base *big.Rat
	exps []*big.Rat
}

var (
	oneHalf *exponentCache = newExponentCache(big.NewRat(1,2))
	oneThird *exponentCache = newExponentCache(big.NewRat(1,3))
	twoThirds *exponentCache = newExponentCache(big.NewRat(2,3))
)

func (ec *exponentCache) addEntry() {
	if ec.exps == nil {
		ec.exps = make([]*big.Rat, 1)
		ec.exps[0] = big.NewRat(1,1)
		return
	}
	newEntry := big.NewRat(1,1)
	newEntry.Mul(ec.base, ec.exps[len(ec.exps)-1])
	ec.exps = append(ec.exps, newEntry)
}

func (ec *exponentCache) toThe(n uint) *big.Rat {
	for n >= uint(len(ec.exps)) {
		ec.addEntry()
	}
	return ec.exps[n]
}

func newExponentCache(base *big.Rat) *exponentCache {
	return &exponentCache{base: base, exps: nil}
}
