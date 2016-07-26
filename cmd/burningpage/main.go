package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/TheDorkKnight/burningdice"
)

func alert(msg string) {
	js.Global.Call("alert", msg)
}

func getElementById(elem string) *js.Object {
	return js.Global.Get("document").Call("getElementById", elem)
}

func usage(program string) string {
	u := "usage: "
	u += program
	u += " [Ob] [NumDice]\n"
	u += "\t     Ob: The number of successes (4-6) required\n"
	u += "\tNumDice: The number of dice to roll\n"
	return u
}

func calculateChances() {
	target := getElementById("ob").Get("value").Uint64()
	numDice := getElementById("num-dice").Get("value").Uint64()
	
	nonOpenChance := burningdice.ChanceOfSuccessesOrGreater(uint(target), uint(numDice))
	nonOpenPercent, _ := nonOpenChance.Float32()
	nonOpenPercent *= 100
	openChance := burningdice.ChanceOfSuccessesOrGreaterWithOpenSixes(uint(target), uint(numDice))
	openPercent, _ := openChance.Float32()
	openPercent *= 100

	getElementById("natural").Set("innerHTML", fmt.Sprintf("Chance of rolling %d or more successes with %d dice: %v (%.2f%%)\n", target, numDice, nonOpenChance, nonOpenPercent))
	getElementById("open").Set("innerHTML", fmt.Sprintf("Chance of rolling %d or more successes with %d dice and open 6's: %v (%.2f%%)\n", target, numDice, openChance, openPercent))
}

func main() {
	js.Global.Set("burningdice", map[string]interface{}{
		"calculate" : calculateChances,
	})
}
