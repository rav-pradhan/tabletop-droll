package main

import (
	"math/rand"
	"time"

	"github.com/rav-pradhan/droll/service"
)

var rolls []string
var results []service.RollResult

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	service.SetFlags(&rolls)
	results = service.ParseRollValues(rolls)
	service.Format(results)
}
