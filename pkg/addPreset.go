package addPreset

import (
	"encoding/json"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//This function is made entirely for the preset window.

func AddPresetWindow() {
	//Add preset window
	var presetWindow fyne.Window = fyne.CurrentApp().NewWindow("Add Preset")

	//Add a drop-down with the preset type (frequency by default)
	presetTypeDropdown := widget.NewSelect([]string{"Default (Frequency)"}, func(value string) {
		// Logic to select preset from JSON map later
	})

	//Add a input for the preset name
	presetNameInput := widget.NewEntry()
	presetNameInput.SetPlaceHolder("Preset Name")

	//Add a input for the preset description
	presetDescInput := widget.NewEntry()
	presetDescInput.SetPlaceHolder("Preset Description")

	//Add a input for the preset frequencies
	presetFrequenciesInput := widget.NewEntry()
	presetFrequenciesInput.SetPlaceHolder("Preset Frequencies")

	//Add a ok button
	presetOkButton := widget.NewButton("OK", func() {
		title := presetNameInput.Text
		desc := presetDescInput.Text
		frequencies := presetFrequenciesInput.Text

		//Save type, asked inputs and name on a json, each on a different dictionarie defined by the preset name.
		m := map[string]interface{}{
			title: map[string]interface{}{
				"Type":        presetTypeDropdown.Selected,
				"Description": desc,
				"Frequencies": frequencies,
			},
		}

		//Save preset to json
		file, _ := json.MarshalIndent(m, "", " ")
		_ = os.WriteFile("presets.json", file, 0644)
		presetWindow.Close()

	})

	//Add a cancel button
	presetCancelButton := widget.NewButton("Cancel", func() {
		presetWindow.Close()
	})

	presetLayout := container.NewVBox(
		presetTypeDropdown,
		presetNameInput,
		presetDescInput,
		presetFrequenciesInput,
		container.NewHBox(presetOkButton, presetCancelButton),
	)

	presetWindow.SetContent(presetLayout)
	presetWindow.Resize(fyne.NewSize(500, 300))
	presetWindow.Show()
}
