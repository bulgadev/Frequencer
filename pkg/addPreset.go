package addPreset

import (
	utils "Frequencer/pkg/utils"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Single_wave struct {
	Name string

	Desc string

	Frequencie int
}

type Bineural_beats struct {
	Name string

	Desc string

	FrequencieLeft int

	FrequencieRight int

	Delta int
}

func generateFormFromStruct(data interface{}, entryMap map[string]*widget.Entry) *widget.Form {
	form := widget.NewForm()

	v := reflect.ValueOf(data)
	if !v.IsValid() {
		return form
	}

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return form
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		fieldType := t.Field(i)

		if fieldType.PkgPath != "" {
			continue
		}

		label := fieldType.Name

		switch fieldVal.Kind() {
		case reflect.String:
			entry := widget.NewEntry()
			entry.SetText(fmt.Sprintf("%v", fieldVal.Interface()))
			entryMap[label] = entry
			form.Append(label, entry)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			entry := widget.NewEntry()
			entry.SetText(fmt.Sprintf("%d", fieldVal.Int()))
			entryMap[label] = entry
			form.Append(label, entry)
		default:
			entry := widget.NewEntry()
			entry.Disable()
			form.Append(label, entry)
		}
	}

	return form
}

func AddPresetWindow(presetDropdown *widget.Select) {
	var presetWindow fyne.Window = fyne.CurrentApp().NewWindow("Add Preset")

	dynamicFields := container.NewVBox()
	entryMap := make(map[string]*widget.Entry)

	presetTypeDropdown := widget.NewSelect([]string{"Default (Frequency)", "Binaural Beats"}, func(value string) {
		dynamicFields.Objects = nil
		entryMap = make(map[string]*widget.Entry)
		switch value {
		case "Binaural Beats":
			dynamicFields.Add(generateFormFromStruct(Bineural_beats{}, entryMap))
		default:
			dynamicFields.Add(generateFormFromStruct(Single_wave{}, entryMap))
		}
		dynamicFields.Refresh()
	})

	presetOkButton := widget.NewButton("OK", func() {
		preset := map[string]interface{}{
			"Type": presetTypeDropdown.Selected,
		}

		for fieldName, entry := range entryMap {
			preset[fieldName] = entry.Text
		}

		title := entryMap["Name"].Text

		_ = utils.SavePreset(title, preset)
		presetWindow.Close()

		utils.ClearCache()
		utils.CleanPresets(presetDropdown)

		utils.LoadPresets(presetDropdown)
	})

	presetCancelButton := widget.NewButton("Cancel", func() {
		presetWindow.Close()
	})

	presetLayout := container.NewVBox(
		presetTypeDropdown,
		dynamicFields,
		container.NewHBox(presetOkButton, presetCancelButton),
	)

	presetWindow.SetContent(presetLayout)
	presetWindow.Resize(fyne.NewSize(500, 300))
	presetWindow.Show()
}
