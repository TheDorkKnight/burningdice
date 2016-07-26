package burningdice

import (
	"math/big"
)

func ChanceOfSuccessesOrGreaterWithOpenSixes(targetSuccesses, numDice uint) *big.Rat {
	chance := ChanceOfSuccessesOrGreater(targetSuccesses, numDice)
	
	var initialSuccesses uint
	for initialSuccesses = 1; initialSuccesses < targetSuccesses; initialSuccesses++ {
		chanceOfThisManySuccesses := chanceOfSuccesses(initialSuccesses, numDice)
		
		var sixes uint
		for sixes = 1; sixes <= initialSuccesses; sixes++ {
			chanceOfThisManySixes := chanceOfSixesGivenSuccesses(sixes, initialSuccesses)
			chanceOfOpening := ChanceOfSuccessesOrGreaterWithOpenSixes(targetSuccesses - initialSuccesses, sixes)
			chanceOfOpening = chanceOfOpening.Mul(chanceOfOpening, chanceOfThisManySixes)
			chanceOfOpening = chanceOfOpening.Mul(chanceOfOpening, chanceOfThisManySuccesses)
			chance = chance.Add(chance, chanceOfOpening)
		}
	}
	return chance
}

func ChanceOfSuccessesOrGreater(targetSuccesses, numDice uint) *big.Rat {
	chance := big.NewRat(0,1)
	for enoughSuccesses := targetSuccesses; enoughSuccesses <= numDice; enoughSuccesses++ {
		chance = chance.Add(chance, chanceOfSuccesses(enoughSuccesses, numDice))
	}
	return chance
}

type exponentiable interface {
	toThe(n uint) *big.Rat
}

func chanceOfN(targetNumRolls, numDiceTotal uint, chanceForRoll, chanceAgainstRoll exponentiable) *big.Rat {
	if targetNumRolls > numDiceTotal {
		return big.NewRat(0,1)
	}

	successCombinations := combinations(targetNumRolls, numDiceTotal)
	
	chance := big.NewRat(1,1).SetInt(successCombinations)
	chance = chance.Mul(chance, chanceForRoll.toThe(targetNumRolls))
	return chance.Mul(chance, chanceAgainstRoll.toThe(numDiceTotal - targetNumRolls))
}

func chanceOfSuccesses(targetSuccesses, numDice uint) *big.Rat {
	return chanceOfN(targetSuccesses, numDice, oneHalf, oneHalf)
}

func chanceOfSixesGivenSuccesses(targetSixes, numSuccesses uint) *big.Rat {
	return chanceOfN(targetSixes, numSuccesses, oneThird, twoThirds)
}
