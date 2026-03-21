# Frequencer - AI Self FAQ

### How is audio played?
Using `github.com/ebitengine/oto/v3`. The app sets up a context (`SampleRate=44100`, `ChannelCount=2`, `FormatSignedInt16LE`), and relies on `io.Reader` implementations (like `SineWave.Read`) to feed PCM byte data to the audio player.

### How are presets loaded and used?
- They are loaded from `presets.json` into a UI dropdown.
- Selecting one and hitting "Run" retrieves the config from memory using `utils.GetCached(name)`.
- The data is passed to `utils.RunAudio` which routes it to the corresponding wave builder function via `waveModules` map.

### Why do some presets have different fields?
Because different audio types require different parameters. For example, "Default (Frequency)" needs a single `Frequencies` string, while "Binaural Beats" needs `FrequencieLeft`, `FrequencieRight`, and `Delta`. Because of this, the preset model uses a unified struct and provides a `Summary()` method to display only relevant, non-empty fields.

### How to add a new wave type module?
1. Create a new struct that implements `io.Reader` (`Read(p []byte) (n int, err error)`).
2. Create a builder function for it: `func buildMyWave(sampleRate int, info models.Presets) io.Reader`.
3. Register the builder function in the `waveModules` map in `pkg/utils/runAudio.go`.

### How can I stop a frequency from playing?
Currently, the app follows a "one-at-a-time" policy. When you click "Run" for a new preset, `utils.RunAudio` automatically identifies the `activePlayer`, pauses it, and replaces it with the new frequency. This ensures frequencies don't override or stack on each other.
