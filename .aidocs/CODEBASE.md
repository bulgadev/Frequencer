# Frequencer Codebase Structure

## Overview
Frequencer is a Go-based desktop application using the `fyne` GUI library and `ebitengine/oto` for audio generation. It allows users to play generated audio waves (like Sine waves and Binaural beats).

## Key Components

### 1. Main Entry (`main.go`)
- Initializes the Fyne application and window.
- Sets up the UI: dropdown for presets, "Run" button to play audio. 
- Selecting a preset and clicking "Run" retrieves the data from cache and updates the UI using `info.Summary()`.
- Triggers `utils.RunAudio` to start playback.

### 2. Audio Engine (`pkg/utils/runAudio.go`)
- Contains the core audio generation logic using `github.com/ebitengine/oto/v3`.
- Defines an interface/structs for different wave types (e.g., `SineWave`, `BinauralBeat`).
- `waveModules`: A map of builder functions to instantiate different wave generators based on the `"Type"` of the preset.
- `RunAudio`: Dispatches the audio stream to the `oto` player. It tracks an `activePlayer` and ensures only one frequency plays at a time by pausing the old player before starting the new one.

### 3. Presets, Cache, and Persistence
- **Models (`pkg/models/presets.go`)**: Defines the data structure for a Preset and includes a `Summary()` method for dynamic information display.
- **Cache (`pkg/utils/cacheHandler.go`)**: In-memory cache holding loaded presets mapping a preset name to a `models.Presets` object.
- **Data Source (`presets.json`)**: JSON file storing the defined presets parameters.
- **Persistence (`pkg/utils/jsonHandler.go` & `pkg/utils/loadPreset.go`)**: Handles reading/writing presets to `presets.json`.

### 4. UI Modals (`pkg/addPreset.go`)
- `AddPresetWindow`: A dynamic form generator that uses reflection to create input fields based on wave type structs (`Single_wave`, `Bineural_beats`) for adding new presets.

## To-Do / Needs Implementation
- Binaural Beats core audio math (currently a stub with ghost variables).
- Dynamic parameter passing from Preset structs to the audio wave builders.
