package models

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
