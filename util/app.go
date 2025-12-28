package util

import (
	"encoding/json"
	"log"
	"os"
)

type App struct {
	Reasons map[string][]string
}

func (app *App) LoadReasons() {
	bytes, err := os.ReadFile("/no.json")

	if err != nil {
		log.Fatal("Could not load no.json file")
	}

	var data map[string][]string

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		log.Fatal("Could not parse no.json file")
	}

	app.Reasons = data
}
