package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	// "time"
)

type symbolFlag struct {
	value string
	set bool
}

func (s *symbolFlag) String() string {
	return s.value
}

func (s *symbolFlag) Set(val string) error {
	s.set = true
	if val == "" {
		s.value = "default" 
	} else {
		s.value = val
	}
	return nil
}

func main() {

	// Defining Character Sets
	upperSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerSet := "abcdefghijklmnopqrstuvwxyz"
	numberSet := "0123456789"
	symbolSet := "!@#$%^&*()"

    preprocessArgs()

	// Reading in Arguments
	length := flag.Int("l", 24, "Password length")
	upper := flag.Bool("u", false, "Include uppercase letters")
	lower := flag.Bool("x", false, "Include lowercase letters")
	number := flag.Bool("n", false, "Include numbers")
	// symbols := flag.Bool("s", false, "Include symbols")
	var sFlag symbolFlag 
	flag.Var(&sFlag, "s", "Custom symbols to include (optional)")
	minType := flag.Int("t", 2, "Minimum occurances of each type selected (i.e. 3 upper, 3 lower)")
	// allowedSymbols := flag.String("S", "default", "Custom symbols to include")
	flag.Parse()

	var specialSymbols string = symbolSet
	symbols := false

	if sFlag.set {
		if sFlag.value != "default" && (!*upper && !*lower && !*number) {
			*upper = true
			*lower = true
			*number = true
			symbols = true
			specialSymbols = sFlag.value
		} else if sFlag.value != "default" {
			specialSymbols = sFlag.value
			symbols = true
		} else {
			symbols = true
		}
	}

	// // if custom symbols are passed check if any other flags were passed -u -x -n and if so ensure -s is true
	// if (*upper || *lower || *number) && *symbols {

	// }

	// if nothing is passed then use all character sets
	if !*upper && !*lower && !*number && !symbols {
		*upper = true
		*lower = true
		*number = true
		symbols = true
	}

	// Creating map to remove duplicate symbols to preserve the odds each avalible char can be selected
	symbolMap := make(map[rune]bool)
	for _, r := range specialSymbols {
	    symbolMap[r] = true
	}

	// Building the symbol string from the map created above
	dedupedSymbols := ""
	for r := range symbolMap {
	    dedupedSymbols += string(r)
	}
	fmt.Println(dedupedSymbols)

	// Setting Character Sets	
	var charset string
	if *upper { charset += upperSet }
	if *lower { charset += lowerSet	}
	if *number { charset += numberSet}
	if symbols { charset += dedupedSymbols	}

	// if charset == "" {
	// 	fmt.Println("No character sets selected. Enable at least one.")
	// 	os.Exit(1)
	// }

	// If no arguments were entered using all character sets
	// if !*upper && !*lower && !*number && !*symbols {
	// 	charset += upperSet + lowerSet + numberSet + symbolSet
	// }

	// Checking if complexity requirements can be meant with the set password length
	var enabledTypes int = 0
	if *upper { enabledTypes++ }
	if *lower { enabledTypes++ }
	if *number { enabledTypes++ }
	if symbols { enabledTypes++ }

	if *minType * enabledTypes > *length {
		fmt.Fprintf(os.Stderr, "Error: Can not satisfy required complexity of %d occurances of %d types in a %d character long password.\nIncrease the password length with -l or reduce complexity with -t\n", *minType, enabledTypes, *length)
		os.Exit(1)
	}

	// Create password slice the length of the password
	password := make([]byte, *length)
	// Track assigned index's
	assignedIndices := make(map[int]bool)
	
	// pick random indices in password where complexity is gaurnteed
	pickIndices := func(count int) []int {
		indices := []int{}
		for len(indices) < count {
			idx := rand.Intn(*length)
			if !assignedIndices[idx] {
				assignedIndices[idx] = true
				indices = append(indices, idx)
			}
		}
		return indices
	}

	// Fill guaranteed characters for each type at the random indices
	fillType := func(charset string) {
		indices := pickIndices(*minType)
		for _, i := range indices {
			password[i] = charset[rand.Intn(len(charset))]
		}
	}

	if *upper { fillType(upperSet) }
	if *lower { fillType(lowerSet) }
	if *number { fillType(numberSet) }
	if symbols { fillType(dedupedSymbols) }

	// Randomly fill and complete password
	for i := 0; i < *length; i++ {
		if password[i] == 0 {
			password[i] = charset[rand.Intn(len(charset))]
		}
	}

	// Print password
	fmt.Println(string(password))

	// rand.Seed(time.Now().UnixNano())
	// password := make([]byte, *length)
	// for i := range password {
	// 	password[i] = charset[rand.Intn(len(charset))]
	// }
	// fmt.Println(string(password))
}

func preprocessArgs() {
	var newArgs []string

	// First, expand combined short flags
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
			for _, ch := range arg[1:] {
				newArgs = append(newArgs, "-"+string(ch))
			}
		} else {
			newArgs = append(newArgs, arg)
		}
	}

	// Then, inject "" after any -s with no explicit value
	finalArgs := []string{os.Args[0]}
	for i := 0; i < len(newArgs); i++ {
		arg := newArgs[i]
		finalArgs = append(finalArgs, arg)

		if arg == "-s" {
			// If next arg is missing or is another flag, inject empty string
			if i+1 >= len(newArgs) || strings.HasPrefix(newArgs[i+1], "-") {
				finalArgs = append(finalArgs, "")
			}
		}
	}

	os.Args = finalArgs
	fmt.Println(os.Args)
}