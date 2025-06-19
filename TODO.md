# rPass TODO

## Core Function
[X] error handling
[] test cross platform (rocky box and windows box)
[X] ensure best random algorithm is being used
[X] ensure a complexity is meant i.e. min of 2 of each type

## UX
[X] -s is required is -S "!@#" this needs polished out
[] update symbol charset with supported symbols that are known to not cause issues
[] document the behaviour of the -s when custom symbols are provided vs not provided
[] add a failure state if arguments are provided that are not supported i.e. rpass uxc should not still generate a password

## Bugs
[X] Fix Symbol bug were it doesnt respect the symbols provided by -S
[X] Fix bug were if -s is passed with values that the password should contain upper lower and numbers as well

## Polish
[] Spell Check
[] Check debugging statements / remove
[] Remove comments that are from just learning the langauge
[] Create functions were it makes sense and clean up main
[] Look up Go Style Guide if one exists

## Completion
[] compile and get instructions on downloading / running
[] update readme
  [] add "why this project" section
[] rename to lowercase r in rPass - looks cooler