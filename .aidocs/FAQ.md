# Frequencer - AI Self FAQ

### How is audio played?
Using `github.com/ebitengine/oto/v3`. The app sets up an audio context with a sample rate of 44100Hz and 2 channels (stereo). It uses `io.Reader` implementations (like `SineWave.Read` or `BinauralBeat.Read`) to push 16-bit PCM byte data to the audio player.

### How are presets loaded and used?
- At startup, `LoadPresets` reads `presets.json` and populates the `presetCache`.
- Selecting a preset in the UI and clicking "Run" fetches the `Presets` struct from cache.
- The entire `Presets` struct is passed to `RunAudio`, which uses the `Type` field to look up the correct generator in `waveModules`.

### How does Binaural Beat generation work?
The `BinauralBeat` generator plays a base frequency in the left ear and a slightly different frequency (Base + Delta) in the right ear. This difference (Delta) creates the perceived "beat" frequency in the brain.

### How are new presets added?
The "Add Preset" window uses reflection on internal structs (`Single_wave`, `Bineural_beats`) to build a form. When "OK" is clicked:
1. Data is collected from the form.
2. `SavePreset` appends the data to `presets.json`.
3. The cache is cleared and reloaded.
4. The main UI dropdown is updated.

### Can I play multiple frequencies at once?
No, the application enforces a single-player policy. `RunAudio` checks if an `activePlayer` exists, pauses it, and replaces it with the new one. This prevents overlapping audio and ensures a clean listening experience.

### Why use reflection for the "Add Preset" form?
Reflection allows the UI to automatically adapt to new wave types and their specific parameters without requiring manual UI code updates for every new feature.

### How do I set the application icon?
The icon is set in `main.go`. It loads `icon.png` using `fyne.LoadResourceFromPath` and then calls `a.SetIcon(resource)` to set it globally for the application. For the executable icon, the packaging tool uses the `--icon` flag.

### How do I build a release version?
For Windows:
```bash
go run fyne.io/tools/cmd/fyne@latest package --os windows --icon icon.png --app-id com.bulgadev.frequencer
```
This requires `FyneApp.toml` to be present for metadata.

### Why does `fyne-cross` fail for macOS?
macOS cross-compilation requires the macOS SDK, which is not packaged with `fyne-cross` due to licensing. You must provide the SDK path or use a macOS-based builder (like GitHub Actions macos-latest runner).

### How to use GitHub Actions for releases?
Create a `.github/workflows/release.yml` using the `fyne-io/fyne-action`. This is the most reliable way to get Linux and macOS binaries without local environment issues.
