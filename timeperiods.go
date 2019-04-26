package aethersetgen

import (
	"math/rand"
)

// TimePeriod is a time period
type TimePeriod struct {
	Name            string
	Characteristics []string
}

func getRandomTimePeriod() TimePeriod {
	periods := getAllTimePeriods()

	return periods[rand.Intn(len(periods)-1)]
}

func getAllTimePeriods() []TimePeriod {
	timePeriods := []TimePeriod{
		TimePeriod{
			Name:            "Old Stone Age",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "New Stone Age",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Bronze Age",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Iron Age",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Persian Empire",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Hellenism",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Roman Empire",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Byzantine Empire",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Dark Ages",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Medieval Ages",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Age of Exploration",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Rennaissance",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Colonialism",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Industrial Revolution",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Victorian Era",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Turn of the Century",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "World War I",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Prohibition/Depression",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "World War II",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Rock and Roll Era",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "21st Century",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Post-Modern",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Near Future",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Future",
			Characteristics: []string{},
		},
		TimePeriod{
			Name:            "Distant Future",
			Characteristics: []string{},
		},
	}

	return timePeriods
}
