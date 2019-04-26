package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/world/pkg/random"
	"github.com/silvergryphongames/aethersetgen"
)

func getSetting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var newSetting aethersetgen.Setting

	random.SeedFromString(id)

	newSetting = aethersetgen.Generate()

	json.NewEncoder(w).Encode(newSetting)
}

func getSettingRandom(w http.ResponseWriter, r *http.Request) {
	var newSetting aethersetgen.Setting

	rand.Seed(time.Now().UnixNano())

	newSetting = aethersetgen.Generate()

	json.NewEncoder(w).Encode(newSetting)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", getSettingRandom)
	r.Get("/{id}", getSetting)

	fmt.Println("Setting Generator API is online.")
	log.Fatal(http.ListenAndServe(":9797", r))
}
