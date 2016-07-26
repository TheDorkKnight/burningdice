package main

import (
	"fmt"
	"github.com/TheDorkKnight/burningdice"
	"os"
	"strconv"
)

func usage(program string) string {
	u := "usage: "
	u += program
	u += " [Ob] [NumDice]\n"
	u += "\t     Ob: The number of successes (4-6) required\n"
	u += "\tNumDice: The number of dice to roll\n"
	return u
}

func main() {
	if (len(os.Args) != 3) {
		fmt.Print(usage(os.Args[0]))
		os.Exit(1)
	}
	target, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Could not parse [Ob]:", err)
		fmt.Print(usage(os.Args[0]))
		os.Exit(2)
	}
	numDice, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Could not parse [NumDice]:", err)
		fmt.Print(usage(os.Args[0]))
		os.Exit(2)
	}
	
	nonOpenChance := burningdice.ChanceOfSuccessesOrGreater(uint(target), uint(numDice))
	nonOpenPercent, _ := nonOpenChance.Float32()
	nonOpenPercent *= 100
	openChance := burningdice.ChanceOfSuccessesOrGreaterWithOpenSixes(uint(target), uint(numDice))
	openPercent, _ := openChance.Float32()
	openPercent *= 100
	
	fmt.Printf("Chance of rolling %d or more successes with %d dice: %v (%.2f%%)\n", target, numDice, nonOpenChance, nonOpenPercent) 
	fmt.Printf("Chance of rolling %d or more successes with %d dice and open 6's: %v (%.2f%%)\n", target, numDice, openChance, openPercent)
}
