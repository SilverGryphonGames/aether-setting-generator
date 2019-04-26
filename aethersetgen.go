package aethersetgen

// Setting is a conglomeration of descriptors for a game setting
type Setting struct {
	Description string
	TimePeriod  TimePeriod
	Technology  Technology
	Tone        Tone
	Genres      []Genre
}

// Generate generates a random setting
func Generate() Setting {
	setting := Setting{}
	setting.TimePeriod = getRandomTimePeriod()
	setting.Technology = getRandomTechnology()
	setting.Tone = getRandomTone()
	setting.Genres = getRandomGenres()
	setting.Description = setting.getDescription()

	return setting
}

func (setting Setting) getDescription() string {
	var description string

	description = "In a " + setting.TimePeriod.Name + " time period, using " + setting.Technology.Name + " technology. The tone is " + setting.Tone.Name + ", and the setting fits the " + setting.Genres[0].Name + " and " + setting.Genres[1].Name + " genres."

	return description
}
