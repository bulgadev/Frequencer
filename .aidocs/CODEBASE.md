# Frequencer Codebase Structure

## Overview
Frequencer is a Go-based desktop application using the `fyne` GUI library and `ebitengine/oto` for audio generation. It allows users to play generated audio waves (like Sine waves and Binaural beats).

## Key Components

### 1. Main Entry (`main.go`)
- Initializes the Fyne application and window.
- Sets up the UI: dropdown for presets, "Run" button to play audio.
- Retreives preset data from the cache (via `utils.GetCached`) and triggers `utils.RunAudio`.

### 2. Audio Engine (`pkg/utils/runAudio.go`)
- Contains the core audio generation logic using `github.com/ebitengine/oto/v3`.
- Defines an interface/structs for different wave types (e.g., `SineWave`, `BinauralBeat`).
- `waveModules`: A map of builder functions to instantiate different wave generators based on the `"Type"` of the preset.
- `RunAudio`: Dispatches the audio stream to the `oto` player.

### 3. Presets and Cache 
- **Models (`pkg/models/presets.go`)**: Defines the data structure for a Preset.
- **Cache (`pkg/utils/cacheHandler.go`)**: In-memory cache holding loaded presets mapping a preset name to a `models.Presets` object.
- **Data Source (`presets.json`)**: JSON file storing the defined presets parameters. Includes fields like Description, Type, Frequencies (for Default), or FrequencieLeft, FrequencieRight, Delta (for Binaural).

## To-Do / Needs Implementation
- Binaural Beats core audio math (currently a stub with ghost variables).
- Dynamic parameter passing from Preset structs to the audio wave builders.
