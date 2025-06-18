package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	symbols := flag.Bool("s", false, "Include symbols")
	allowedSymbols := flag.String("S", symbolSet, "Custom symbols to include")
	flag.Parse()

	// Creating map to remove duplicate symbols to preserve the odds each avalible char can be selected
	symbolMap := make(map[rune]bool)
	for _, r := range *allowedSymbols {
	    symbolMap[r] = true
	}

	// Building the symbol string from the map created above
	dedupedSymbols := ""
	for r := range symbolMap {
	    dedupedSymbols += string(r)
	}
	fmt.Println(dedupedSymbols)

	// Checking if arguments were provided and mapping those character sets	
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
	if *symbols {
		charset += dedupedSymbols
	}

	// if charset == "" {
	// 	fmt.Println("No character sets selected. Enable at least one.")
	// 	os.Exit(1)
	// }

	// If no arguments were entered using all character sets
	if !*upper && !*lower && !*number && !*symbols {
		charset += upperSet + lowerSet + numberSet + symbolSet
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, *length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Println(string(password))
}

func preprocessArgs() {   
	var newArgs []string
	// Iterate over OS Args excluding 0 as that is the program name
    for _, arg := range os.Args[1:] {
		// Checking if args were combined (i.e. -unx)
        if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
            // Iterate through each arg in the combined argument and seperate them
            for _, ch := range arg[1:] {
                newArgs = append(newArgs, "-"+string(ch))
            }
		} else {
			newArgs = append(newArgs, arg)
		}
	}
	// Remake the arg list starting with the os name and each argument indivdually (i.e. rpass -u -x -n)
	os.Args = append([]string{os.Args[0]}, newArgs...)
	fmt.Println(os.Args)
}