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
	//tempWarn := widget.NewLabel(fmt.Sprintf("This is the first stable application version. %s The actual sound system is not implemented yet. %s So if you came for the code, its fine, but the app is not ready yet.", "\n", "\n"))

	infoText := widget.NewLabel("")
	infoText.Wrapping = fyne.TextWrapWord

	///
	//Add a drop-down with the preset type (frequency by default)
	///
	presetDropdown := widget.NewSelect([]string{"None"}, func(value string) {
		//Logic to select preset from JSON map later
		name = value
	})

	///
	//Get from cache, and run it
	///
	testButton := widget.NewButton("Run", func() {
		if name != "None" && name != "" {
			info, _ := utils.GetCached(name)
			textQuery := fmt.Sprintf("Name: %s\nDescription: %s\nFrequencies: %s\nType: %s", name, info.Description, info.Frequencies, info.Type)
			infoText.SetText(textQuery)
			// Refactored to pass the entire struct for dynamic parameter fetching
			utils.RunAudio(info)
			errorText.Hide()
		} else {
			errorText.Show()
			errorText.SetText("No preset selected")
		}
	})

	///
	//Add a button to add new preset
	///
	addPresetButton := fyne.NewMenuItem("Add Preset", func() {
		addPreset.AddPresetWindow(presetDropdown)
	})

	//Load presets from json
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
			//tempWarn,
			errorText,
		),
	)

	w.SetOnClosed(func() {
		utils.ClearCache()
		utils.CleanPresets(presetDropdown)
	})

	w.Resize(fyne.NewSize(500, 150))
	w.ShowAndRun()
}
