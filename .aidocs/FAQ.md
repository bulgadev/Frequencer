# Frequencer - AI Self FAQ

### How is audio played?
Using `github.com/ebitengine/oto/v3`. The app sets up a context (`SampleRate=44100`, `ChannelCount=2`, `FormatSignedInt16LE`), and relies on `io.Reader` implementations (like `SineWave.Read`) to feed PCM byte data to the audio player.

### How are presets loaded and used?
- They are loaded from `presets.json` into a UI dropdown.
- Selecting one and hitting "Run" retrieves the config from memory using `utils.GetCached(name)`.
- The data is passed to `utils.RunAudio` which routes it to the corresponding wave builder function via `waveModules` map.

### Why do some presets have different fields?
Because different audio types require different parameters. For example, "Default (Frequency)" needs a single `Frequencies` string, while "Binaural Beats" needs `FrequencieLeft`, `FrequencieRight`, and `Delta`. Because of this, the preset model uses a unified struct that might have empty fields depending on the preset type.

### How to add a new wave type module?
1. Create a new struct that implements `io.Reader` (`Read(p []byte) (n int, err error)`).
2. Create a builder function for it: `func buildMyWave(sampleRate int, info models.Presets) io.Reader`.
3. Register the builder function in the `waveModules` map in `pkg/utils/runAudio.go`.
