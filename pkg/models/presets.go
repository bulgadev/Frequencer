package models

import "fmt"

type Presets struct {
	Name            string `json:"Name"`
	Description     string `json:"Description,omitempty"`
	Desc            string `json:"Desc,omitempty"`           // Handles alternative descriptor string dynamically
	Frequencies     string `json:"Frequencies,omitempty"`    // Used by Default
	FrequencieLeft  string `json:"FrequencieLeft,omitempty"` // Added for Binaural beats Left channel
	FrequencieRight string `json:"FrequencieRight,omitempty"`// Added for Binaural beats Right channel
	Delta           string `json:"Delta,omitempty"`          // Added for Binaural beats beatFreq
	Type            string `json:"Type"`
}

// Summary returns a formatted string containing all non-empty fields.
func (p Presets) Summary() string {
	res := fmt.Sprintf("Name: %s\nType: %s", p.Name, p.Type)
	
	if p.Description != "" {
		res += fmt.Sprintf("\nDescription: %s", p.Description)
	} else if p.Desc != "" {
		res += fmt.Sprintf("\nDescription: %s", p.Desc)
	}

	if p.Frequencies != "" {
		res += fmt.Sprintf("\nFrequencies: %s", p.Frequencies)
	}
	if p.FrequencieLeft != "" {
		res += fmt.Sprintf("\nFrequency Left: %s", p.FrequencieLeft)
	}
	if p.FrequencieRight != "" {
		res += fmt.Sprintf("\nFrequency Right: %s", p.FrequencieRight)
	}
	if p.Delta != "" {
		res += fmt.Sprintf("\nDelta: %s", p.Delta)
	}

	return res
}
