package service

import "github.com/spf13/pflag"

// SetFlags will set the CLI flags and parse values
func SetFlags(rolls *[]string) {
	pflag.StringSliceVar(rolls, "roll", []string{"1d6"}, "Specify the number/sides of die/dice you wish to roll")
	pflag.Parse()
}
