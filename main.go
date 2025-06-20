package main

import (
	"flag"
	"fmt"

	// "math/rand"
	"crypto/rand"
	"math/big"
	"os"
	"strings"
	// "time"
)

type symbolFlag struct {
	value string
	set   bool
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

func secureRandChar(charset string) (byte, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, fmt.Errorf("secureRandChar: %w", err)
	}
	return charset[n.Int64()], nil
}

func secureRandInt(max int) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("secureRandChar: %w", err)
	}
	return n.Int64(), nil
}

func preprocessArgs() {
	/* This function is to allow for complex arguments to be passed -uxn or -uxl 12
	   To do this is splits the arguments up and then reassigns them prior to calling parse()
	*/
	var newArgs []string

	// Expand any arguments and pass in short form i.e. (-uxn becomes -u -x -n)
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
			for _, ch := range arg[1:] {
				newArgs = append(newArgs, "-"+string(ch))
			}
		} else {
			newArgs = append(newArgs, arg)
		}
	}

	// If -s is passed without any special characters add "" else keep speical characters
	finalArgs := []string{os.Args[0]}
	for i := 0; i < len(newArgs); i++ {
		arg := newArgs[i]
		finalArgs = append(finalArgs, arg)
		if arg == "-s" {
			if i+1 >= len(newArgs) || strings.HasPrefix(newArgs[i+1], "-") {
				finalArgs = append(finalArgs, "")
			}
		}
	}
	os.Args = finalArgs
}

func main() {
	const version = "1.0.0"
	upperSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerSet := "abcdefghijklmnopqrstuvwxyz"
	numberSet := "0123456789"
	symbolSet := "!@#$%^&*()"

	// Need struct for -s as it can ether be passed alone -s for default or -s "!@#$" to denote symbols to use
	var sFlag symbolFlag

	length := flag.Int("l", 24, "Password length")
	upper := flag.Bool("u", false, "Include uppercase letters")
	lower := flag.Bool("x", false, "Include lowercase letters")
	number := flag.Bool("n", false, "Include numbers")
	flag.Var(&sFlag, "s", "Custom symbols to include (optional)")
	minType := flag.Int("t", 2, "Minimum occurrences of each type selected (i.e. 3 upper, 3 lower)")
	checkVersion := flag.Bool("version", false, "Display the current version of rPass")

	preprocessArgs()
	flag.Parse()

	if *checkVersion {
		fmt.Println("rPass version", version)
		os.Exit(0)
	}

	var specialSymbols string = symbolSet
	symbols := false

	/* Check if -s was passed
	   If passed alone and with special characters, ensure that upper, lower, and numbers are still used
	   If passed with other arguments and with special characters use the special characters
	   If passed with no special characters set True and use default symbolSet
	*/
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

	// If nothing is passed then use all character sets
	if !*upper && !*lower && !*number && !symbols {
		*upper = true
		*lower = true
		*number = true
		symbols = true
	}

	// Remove Duplicate Symbols that may exist to preserve randomness
	symbolMap := make(map[rune]bool)
	for _, r := range specialSymbols {
		symbolMap[r] = true
	}
	dedupedSymbols := ""
	for r := range symbolMap {
		dedupedSymbols += string(r)
	}

	// Set characters for random selection
	var charset string
	if *upper {
		charset += upperSet
	}
	if *lower {
		charset += lowerSet
	}
	if *number {
		charset += numberSet
	}
	if symbols {
		charset += dedupedSymbols
	}

	// Checking if complexity requirements can be meant with the set password length
	var enabledTypes int = 0
	if *upper {
		enabledTypes++
	}
	if *lower {
		enabledTypes++
	}
	if *number {
		enabledTypes++
	}
	if symbols {
		enabledTypes++
	}

	if *minType*enabledTypes > *length {
		fmt.Fprintf(os.Stderr, "Error: Can not satisfy required complexity of %d occurrences of %d types in a %d character long password.\nIncrease the password length with -l or reduce complexity with -t\n", *minType, enabledTypes, *length)
		os.Exit(1)
	}

	// Create password slice the length of the password
	password := make([]byte, *length)
	// Track assigned indices
	assignedIndices := make(map[int]bool)

	// Select random indices were complexity is guaranteed. This ensures that the -t is enforced in the random password.
	pickIndices := func(count int) ([]int, error) {
		indices := []int{}
		for len(indices) < count {
			idx64, err := secureRandInt(*length)
			if err != nil {
				return nil, fmt.Errorf("pickIndices: %w", err)
			}
			idx := int(idx64)
			if !assignedIndices[idx] {
				assignedIndices[idx] = true
				indices = append(indices, idx)
			}
		}
		return indices, nil
	}

	// Fill selected indexs with one of the enabled types to meet complexity
	fillType := func(charset string) error {
		indices, err := pickIndices(*minType)
		if err != nil {
			return fmt.Errorf("fillType: %w", err)
		}
		for _, i := range indices {
			password[i], err = secureRandChar(charset)
			if err != nil {
				return fmt.Errorf("fillType: %w", err)
			}
		}
		return nil
	}

	if *upper {
		err := fillType(upperSet)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if *lower {
		err := fillType(lowerSet)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if *number {
		err := fillType(numberSet)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if symbols {
		err := fillType(dedupedSymbols)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Randomly fill and complete password
	for i := 0; i < *length; i++ {
		if password[i] == 0 {
			var err error
			password[i], err = secureRandChar(charset)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}

	// Print password to terminal
	fmt.Println(string(password))
}
