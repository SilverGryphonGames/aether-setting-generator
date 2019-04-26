package aethersetgen

import (
	"math/rand"
)

// Genre is a genre
type Genre struct {
	Name            string
	Characteristics []string
}

func getRandomGenre() Genre {
	genres := getAllGenres()

	return genres[rand.Intn(len(genres)-1)]
}

func getRandomGenres() []Genre {
	var genres []Genre
	var genre Genre

	numGenres := 0

	for numGenres < 2 {
		genre = getRandomGenre()
		if !isGenreInSlice(genre, genres) {
			genres = append(genres, genre)
			numGenres++
		}
	}

	return genres
}

func getAllGenres() []Genre {
	genres := []Genre{
		Genre{
			Name:            "Alternate History",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Classic",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Comedy",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Epic",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Fantasy",
			Characteristics: []string{},
		},
		Genre{
			Name:            "High Fantasy",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Low Fantasy",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Horror",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Mystery",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Opera",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Punk",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Romance",
			Characteristics: []string{},
		},
		Genre{
			Name:            "Science Fiction",
			Characteristics: []string{},
		},
	}

	return genres
}

func isGenreInSlice(genre Genre, genres []Genre) bool {
	for _, g := range genres {
		if g.Name == genre.Name {
			return true
		}
	}
	return false
}
