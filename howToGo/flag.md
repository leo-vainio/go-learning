# Command-line flags in Go

### Example runs:

```console

> go build -o out
> out 
> out -bool=true
> out --port 1234 -num=44 --bool -svar=hello
> out --bool=true coffee
> out --bool=true coffee -cream --sugar=brown randomstring
> out --port 1234 -num=44 --bool -svar=hello tea -milk=TRUE

```

```go

// flag.go
package main

import (
	"flag"
	"fmt"
)

// We can bind the flag to a variable using the Var() function. When you bind to variables the variable contain
// the flags value and not a pointer like when you use the flag directly.
var svar string

func init() {
	flag.StringVar(&svar, "svar", "Hi there!", "A string variable")
}

func main() {

	// Define and initialize a flag.
	// flagPtr := flag.<Type>("<identifier>", "<defaultValue>", "<help message>")
	// flag.<Type>() returns a pointer to the value.
	strPtr := flag.String("port", ":8080", "a string")
	intPtr := flag.Int("num", 22, "an int")
	boolPtr := flag.Bool("bool", false, "a bool")

	// After all flags are defined, call flag.Parse() to parse the command line into the defined flags.
	// Parse() parses the command-line flags from os.Args[1:]
	flag.Parse()
	fmt.Println("The flags have been parsed:", flag.Parsed())

	fmt.Println("port:", *strPtr)
	fmt.Println("num:", *intPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println(flag.NArg())  // Number of arguments remaining after flags have been processed.
	fmt.Println(flag.NFlag()) // Number of command-line flags that have been set.

	// Changing the value of a flag.
	err := flag.Set("port", "8181")
	if err != nil {
		return
	}
	fmt.Println("port:", *strPtr)

	// Get the Flag given the name of the flag.
	fmt.Println(flag.Lookup("port"))

	// After parsing, the arguments following the flags are available as the slice flag.Args()
	// or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1.
	// Example: out -port=1234 55 66 77 -> flag.Args() = [55, 66, 77]
	// The command-line arguments need to be specified after the flags.
	fmt.Println("tail:", flag.Args())

	// FlagSet:
	// The flagset uses the remaining arguments after the flags.
	coffee := flag.NewFlagSet("coffee", flag.ExitOnError)
	cream := coffee.Bool("cream", false, "Cream")
	sugar := coffee.String("sugar", "", "Sugar")

	tea := flag.NewFlagSet("tea", flag.ExitOnError)
	milk := tea.Bool("milk", false, "Milk")

	subArgs := flag.Args()
	switch flag.Arg(0) {
	case "coffee":
		coffee.Parse(subArgs[1:])
		fmt.Println("subcommand 'coffee'")
		fmt.Println("  Cream:", *cream)
		fmt.Println("  Sugar:", *sugar)
		fmt.Println("  tail:", coffee.Args())
	case "tea":
		tea.Parse(subArgs[1:])
		fmt.Println("subcommand 'tea'")
		fmt.Println(" Milk:", *milk)
		fmt.Println(" tail:", tea.Args())
	default:
		fmt.Println("No subcommand specified!")
	}

	// Nothing has been removed from the original arguments. The tail of 'coffee' or 'tea' is different than
	// the flag.Args() obviously. So if we wanted we could use these tails to create even more FlagSets. It all depends
	// on how we want the application to run. Too much complexity is harder to maintain. Although, within the switch
	// statements it could be useful since we then know better what is expected for each command used.
	fmt.Println("tail:", flag.Args())
}

```