# burningdice
Library for calculating dice roll chances in Burning Wheel RPG

## func ChanceOfSuccessesOrGreater(targetNumberOfSuccesses, totalNumberOfDice uint) *big.Rat
This function calculates the chance of rolling the given number of successes without opening up 6's

## func ChanceOfSuccessesOrGreaterWithOpenSixes(targetNumberOfSuccesses, totalNumberOfDice uint) *big.Rat
This function calculates the chance of rolling the given number or successes when you can re-roll 6's for extra successes
