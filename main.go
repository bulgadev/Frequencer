package main

import (
	addPreset "Frequencer/pkg"
	utils "Frequencer/pkg/utils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// open a new window (at min 800x600, name frequenced
func main() {

	a := app.New()
	w := a.NewWindow("Hello")

	var name string

	errorText := widget.NewLabel("")

	infoText := widget.NewLabel("")
	infoText.Wrapping = fyne.TextWrapWord

	//Add a drop-down with the preset type (frequency by default)
	presetDropdown := widget.NewSelect([]string{"None"}, func(value string) {
		//Logic to select preset from JSON map later
		name = value
	})

	//Get info from cache
	testButton := widget.NewButton("Run", func() {
		if name != "None" && name != "" {
			info, _ := utils.GetCached(name)
			textQuery := fmt.Sprintf("Name: %s\nDescription: %s\nFrequencies: %s\nType: %s", name, info.Description, info.Frequencies, info.Type)
			infoText.SetText(textQuery)
			errorText.Hide()
		} else {
			errorText.Show()
			errorText.SetText("No preset selected")
		}
	})

	//Add a button to add new preset
	addPresetButton := fyne.NewMenuItem("Add Preset", func() {
		addPreset.AddPresetWindow(presetDropdown)
	})

	utils.LoadPresets(presetDropdown)

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
			infoText,
			errorText,
		),
	)

	w.SetOnClosed(func() {
		utils.ClearCache()
		utils.CleanPresets(presetDropdown)
	})

	w.Resize(fyne.NewSize(500, 300))
	w.ShowAndRun()
}

//Then, on the main menu, add a drop-down with all presets. To add it just loop at the json, and take each dic name

//Then a start button

//When the start button is clicked, it will have to call some audio framework which will be able to generate this frequencies.

//You will take the selected preset from the drop-down, and search for it on the json (maybe a loop? Idk what's the most efficient way to look up a json, I have to search first)

//Then take each necessary variable from the json, pass it to the audio lib, and run it.
