package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/SilverGryphonGames/aethersetgen"
)

func formatSetting(setting aethersetgen.Setting) string {
	genreString := setting.Genres[0]
	if len(setting.Genres) > 1 {
		genreString += " and " + setting.Genres[1]
	}
	return setting.TimePeriod + ", " + setting.Technology + " tech level " + genreString + " setting with a " + setting.Tone + " tone"
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rootHandler := func(w http.ResponseWriter, req *http.Request) {
		setting := aethersetgen.Generate()
		output := formatSetting(setting)
		io.WriteString(w, output)
	}

	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":9797", nil))
}
