package aethersetgen

import (
	"math/rand"
)

// Technology is a technology level
type Technology struct {
	Name            string
	Characteristics []string
}

func getRandomTechnology() Technology {
	technologies := getAllTechnologies()

	return technologies[rand.Intn(len(technologies)-1)]
}

func getAllTechnologies() []Technology {
	technologies := []Technology{
		Technology{
			Name:            "Technology of the Time",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Advanced",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Weird",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Stone",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Bronze",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Iron",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Gear",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Gunpowder",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Rennaissance",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Steam",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Combustion",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Diesel",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Air",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Fission",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Fusion",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Space",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Cyber",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Genetic",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Gravity",
			Characteristics: []string{},
		},
		Technology{
			Name:            "Dark Energy",
			Characteristics: []string{},
		},
	}

	return technologies
}
