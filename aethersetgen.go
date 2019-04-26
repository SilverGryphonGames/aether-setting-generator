package aethersetgen

// Setting is a conglomeration of descriptors for a game setting
type Setting struct {
	TimePeriod TimePeriod
	Technology Technology
	Tone       Tone
	Genres     []Genre
}

// Generate generates a random setting
func Generate() Setting {
	setting := Setting{}
	setting.TimePeriod = getRandomTimePeriod()
	setting.Technology = getRandomTechnology()
	setting.Tone = getRandomTone()
	setting.Genres = getRandomGenres()

	return setting
}
