package main

import (
	"os"

	"github.com/felixrosatmetlla/mealPlanner/cmd/mealPlanner"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	mealPlanner.Execute()
}
