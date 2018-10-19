package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/SilverGryphonGames/aethersetgen"
)

func displaySetting(setting aethersetgen.Setting) {
	genreString := setting.Genres[0]
	if len(setting.Genres) > 1 {
		genreString += " and " + setting.Genres[1]
	}
	fmt.Println(setting.TimePeriod + ", " + setting.Technology + " tech level " + genreString + " setting with a " + setting.Tone + " tone")
}

func main() {
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")

	flag.Parse()

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	setting := aethersetgen.Generate()

	displaySetting(setting)
}
