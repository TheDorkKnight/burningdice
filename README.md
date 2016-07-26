# burningdice
Library for calculating dice roll chances in Burning Wheel RPG

### Chance Without Open 6's
```go
func ChanceOfSuccessesOrGreater(targetNumberOfSuccesses, totalNumberOfDice uint) *big.Rat
```
This function calculates the chance of rolling the given number of successes without opening up 6's

### Chance With Open 6's
```go
func ChanceOfSuccessesOrGreaterWithOpenSixes(targetNumberOfSuccesses, totalNumberOfDice uint) *big.Rat
```
This function calculates the chance of rolling the given number or successes when you can re-roll 6's for extra successes

### Examples
See [cmd/burningterm](/cmd/burningterm/) and [cmd/burningpage](/cmd/burningpage/)
