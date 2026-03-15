package main

import (
	addPreset "Frequencer/pkg"
	"Frequencer/pkg/models"
	utils "Frequencer/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func loadPresets(presetDropdown *widget.Select) {

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
		utils.CacheIt(name, content)
	}

}

// open a new window (at min 800x600, name frequenced
func main() {

	a := app.New()
	w := a.NewWindow("Hello")

	var name string

	//Add a button to add new preset
	addPresetButton := fyne.NewMenuItem("Add Preset", addPreset.AddPresetWindow)
	testButton := widget.NewButton("Get Cached", func() {
		info, _ := utils.GetCached(name)
		fmt.Println(info.Frequencies)

	})

	//Add a drop-down with the preset type (frequency by default)
	presetDropdown := widget.NewSelect([]string{"None"}, func(value string) {
		//Logic to select preset from JSON map later
		name = value
	})

	loadPresets(presetDropdown)

	mainLayout := container.NewVBox(
		presetDropdown,
		testButton,
	)

	//Add a menu bar with file, to store the add preset button
	mainMenu := fyne.NewMainMenu(fyne.NewMenu("File", addPresetButton))
	w.SetMainMenu(mainMenu)

	//Main window content
	w.SetContent(
		container.NewVBox(
			mainLayout,
		),
	)

	w.Resize(fyne.NewSize(500, 300))
	w.ShowAndRun()
}

//Then, on the main menu, add a drop-down with all presets. To add it just loop at the json, and take each dic name

//Then a start button

//When the start button is clicked, it will have to call some audio framework which will be able to generate this frequencies.

//You will take the selected preset from the drop-down, and search for it on the json (maybe a loop? Idk what's the most efficient way to look up a json, I have to search first)

//Then take each necessary variable from the json, pass it to the audio lib, and run it.
