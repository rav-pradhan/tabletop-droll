package service

import "fmt"

// Format will format the dice results in an easily readable manner
func Format(results []RollResult) {
	for _, r := range results {
		fmt.Printf("\n")
		fmt.Printf("Rolled %q: Score: %d\n", r.rollValue, r.score)
	}
	fmt.Printf("\n")
}
