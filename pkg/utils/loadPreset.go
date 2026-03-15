package utils

import (
	"Frequencer/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2/widget"
)

func LoadPresets(presetDropdown *widget.Select) {

	jsonData, err := os.ReadFile("presets.json")

	if err != nil {
		log.Fatal(err)
	}

	var result map[string]models.Presets

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("error:", err)
	}

	for name, content := range result {
		//Add the name to the drop-down
		presetDropdown.Options = append(presetDropdown.Options, name)

		//Cache the preset
		CacheIt(name, content)
	}

}

func CleanPresets(presetDropdown *widget.Select) {
	presetDropdown.Options = []string{"None"}
}
