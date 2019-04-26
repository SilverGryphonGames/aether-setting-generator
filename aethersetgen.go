package aethersetgen

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Setting is a conglomeration of descriptors for a game setting
type Setting struct {
	TimePeriod string
	Technology string
	Tone       string
	Genres     []string
}

func randomTimePeriod() string {
	return random.String(timePeriods)
}

func randomTechnology() string {
	return random.String(technologies)
}

func randomTone() string {
	return random.String(tones)
}

func randomGenre() string {
	return random.String(genres)
}

func randomGenres() []string {
	var genres []string

	genres = append(genres, randomGenre())

	roll := rand.Intn(10)
	if roll >= 9 {
		genres = append(genres, randomGenre())
	}

	return genres
}

// Generate generates a random setting
func Generate() Setting {
	setting := Setting{}
	setting.TimePeriod = randomTimePeriod()
	setting.Technology = randomTechnology()
	setting.Tone = randomTone()
	setting.Genres = randomGenres()

	return setting
}
