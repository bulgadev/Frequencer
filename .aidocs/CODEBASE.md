# Frequencer Codebase Structure

## Overview
Frequencer is a Go-based desktop application using the `fyne` GUI library and `ebitengine/oto` for audio generation. It allows users to play generated audio waves (like Sine waves and Binaural beats) based on customizable presets.

## Key Components

### 1. Main Entry (`main.go`)
- Initializes the Fyne application and window.
- **Icon Management**: Loads `icon.png` from the root directory and sets it as the global application and window icon.
- Sets up the UI: a dropdown for presets, a "Run" button to play audio, and a "File" menu for adding presets.
- Selecting a preset and clicking "Run" retrieves the preset from cache, updates the UI summary, and triggers `utils.RunAudio`.
- Handles application shutdown by clearing cache and cleaning UI state.

### 2. Audio Engine (`pkg/utils/runAudio.go`)
- Core audio generation logic using `github.com/ebitengine/oto/v3`.
- **Wave Type Registry**: `waveModules` map links wave type names to builder functions.
- **Sine Wave**: Generates a single-frequency sine wave.
- **Binaural Beats**: Generates distinct frequencies for left and right channels to create a binaural beat effect.
- **RunAudio**: Orchestrates playback. It ensures only one frequency plays at a time by pausing any currently active player before starting a new one.

### 3. Presets and Persistence
- **Models (`pkg/models/presets.go`)**: Defines the `Presets` struct which supports both single frequency and binaural parameters (`FrequencieLeft`, `FrequencieRight`, `Delta`). Includes a `Summary()` method for UI display.
- **Cache (`pkg/utils/cacheHandler.go`)**: In-memory storage for loaded presets for quick access during runtime.
- **JSON Handler (`pkg/utils/jsonHandler.go`)**: Handles appending new presets to `presets.json` without overwriting existing ones.
- **Loader (`pkg/utils/loadPreset.go`)**: Reads `presets.json` at startup to populate the UI and cache.

### 4. Dynamic Preset Management (`pkg/addPreset.go`)
- **AddPresetWindow**: A dynamic form generator that uses Go's reflection to create input fields based on the wave type selected (e.g., `Single_wave` vs. `Bineural_beats`).
- Automatically updates the main dropdown and cache after a new preset is saved.

## Implementation Status
- [x] Sine Wave generation.
- [x] Binaural Beats generation (Left frequency + Delta difference).
- [x] Dynamic form generation for adding presets.
- [x] One-at-a-time audio playback policy.
- [x] Persistence via `presets.json`.
- [x] Release build process (Windows).

## Release Process

### Windows Build
To package the application for Windows with the icon and metadata:
```bash
go run fyne.io/tools/cmd/fyne@latest package --os windows --icon icon.png --app-id com.bulgadev.frequencer
```
This generates `Frequencer.exe` in the root directory.

### Cross-Platform (Linux/macOS)
Due to SDK and dependency requirements (especially for macOS), it is recommended to use **GitHub Actions** for cross-platform releases. A template is provided in the `FAQ.md`.
