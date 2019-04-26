package aethersetgen

import (
	"math/rand"
)

// Tone is a tone
type Tone struct {
	Name            string
	Characteristics []string
}

func getRandomTone() Tone {
	tones := getAllTones()

	return tones[rand.Intn(len(tones)-1)]
}

func getAllTones() []Tone {
	tones := []Tone{
		Tone{
			Name:            "Action",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Black",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Dark",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Dystopia",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Heroic",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Noir",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Romance",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Supernatural",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Tragedy",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Utopia",
			Characteristics: []string{},
		},
		Tone{
			Name:            "Comedy",
			Characteristics: []string{},
		},
	}

	return tones
}
