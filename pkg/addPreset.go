package addPreset

import (
	utils "Frequencer/pkg/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//This function is made entirely for the preset window.

func AddPresetWindow(presetDropdown *widget.Select) {
	//Add preset window
	var presetWindow fyne.Window = fyne.CurrentApp().NewWindow("Add Preset")

	//Add a drop-down with the preset type (frequency by default)
	// Wave type options — must match keys in waveModules (runAudio.go)
	presetTypeDropdown := widget.NewSelect([]string{"Default (Frequency)", "Binaural Beats"}, func(value string) {
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
		preset := map[string]interface{}{
			"Type":        presetTypeDropdown.Selected,
			"Description": desc,
			"Frequencies": frequencies,
		}

		//Save preset to json (appends to existing presets)
		_ = utils.SavePreset(title, preset)
		presetWindow.Close()

		//Clear cache and dropdown
		utils.ClearCache()
		utils.CleanPresets(presetDropdown)

		//Add the preset to the dropdown
		utils.LoadPresets(presetDropdown)
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
