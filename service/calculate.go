package service

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
)

// RollResult is the struct which contains the type of roll it was - e.g., 2d8 - and the score after calculation
type RollResult struct {
	rollValue string
	score     int
}

type rollData struct {
	numberOfRolls int
	numberOfSides int
}

// ParseRollValues takes the rolls arguments and parses them into int values for calculating scores
func ParseRollValues(r []string) []RollResult {
	var rollData []RollResult
	for i := 0; i < len(r); i++ {
		rd, err := buildRollData(r[i])
		if err != nil {
			log.Fatal("Error occurred when building roll data: ", err)
		}

		res := RollResult{
			rollValue: r[i],
			score:     calculate(rd.numberOfRolls, rd.numberOfSides),
		}
		rollData = append(rollData, res)
	}
	return rollData
}

func calculate(dieRolls, dieSides int) int {
	var total int
	for i := 1; i <= dieRolls; i++ {
		total += rollDie(dieSides)
	}
	return total
}

func buildRollData(die string) (rollData, error) {
	values := strings.Split(die, "d")

	var err error
	r := rollData{}
	r.numberOfSides, err = strconv.Atoi(values[1])

	if values[0] == "" {
		r.numberOfRolls = 1
	} else {
		r.numberOfRolls, err = strconv.Atoi(values[0])
	}
	return r, err
}

func rollDie(sides int) int {
	min := 1
	return min + rand.Intn(sides-min)
}
