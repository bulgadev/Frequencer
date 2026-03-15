package utils

import (
	"encoding/json"
	"os"
)

const presetsFile = "presets.json"

// SavePreset appends a new preset to the existing presets.json file
// instead of overwriting the entire file.
func SavePreset(name string, preset map[string]interface{}) error {
	existing := make(map[string]interface{})

	// Read existing file if it exists
	data, err := os.ReadFile(presetsFile)
	if err == nil {
		_ = json.Unmarshal(data, &existing)
	}

	// Append (or overwrite) the new preset entry
	existing[name] = preset

	// Write back the full map
	file, err := json.MarshalIndent(existing, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(presetsFile, file, 0644)
}
