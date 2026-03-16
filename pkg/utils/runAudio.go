package utils

import (
	"encoding/binary"
	"io"
	"math"
	"strconv"

	"github.com/ebitengine/oto/v3"
)

// --- Wave type module registry ---
// Maps wave type names (from presets) to their builder functions.
// Each builder returns an io.Reader that generates audio samples.
var waveModules = map[string]func(sampleRate int, freq float64) io.Reader{
	"Default (Frequency)": buildSineWave,
	"Binaural Beats":      buildBinauralBeat,
}

// ========== Default (Frequency) Module ==========

type SineWave struct {
	sampleRate int
	freq       float64
	phase      float64
}

func (s *SineWave) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p)/4; i++ { // 4 bytes = 2 channels * 2 bytes (16-bit)
		// Calculate the sine value
		v := math.Sin(s.phase * 2 * math.Pi)
		s.phase += s.freq / float64(s.sampleRate)
		if s.phase > 1 {
			s.phase -= 1
		}

		// Scale float to 16-bit PCM (max 32767)
		// Using a lower volume (0.3) to avoid clipping
		sample := int16(v * 0.3 * 32767)

		// Write to Left and Right channels (Little Endian)
		binary.LittleEndian.PutUint16(p[i*4:], uint16(sample))
		binary.LittleEndian.PutUint16(p[i*4+2:], uint16(sample))
	}
	return len(p), nil
}

// buildSineWave - builder for the default single sine wave module
func buildSineWave(sampleRate int, freq float64) io.Reader {
	return &SineWave{
		sampleRate: sampleRate,
		freq:       freq,
		phase:      0,
	}
}

// ========== Binaural Beats Module (stub) ==========

// BinauralBeat generates a binaural beat by playing slightly different
// frequencies on the left and right channels.
// GHOST VARIABLES — wire these when implementing the binaural preset:
//   - baseFreq: the carrier/base frequency (Hz) for both ears
//   - beatFreq: the difference frequency (Hz) — left ear gets baseFreq, right ear gets baseFreq + beatFreq
//   - phaseL:   phase accumulator for the left channel
//   - phaseR:   phase accumulator for the right channel
type BinauralBeat struct {
	sampleRate int
	baseFreq   float64 // GHOST: carrier frequency passed from preset
	beatFreq   float64 // GHOST: beat difference frequency passed from preset
	phaseL     float64 // GHOST: left channel phase accumulator
	phaseR     float64 // GHOST: right channel phase accumulator
}

// Read — TODO: implement binaural beat sample generation here
func (b *BinauralBeat) Read(p []byte) (n int, err error) {
	// TODO: generate left channel at baseFreq, right channel at baseFreq + beatFreq
	// Use phaseL / phaseR accumulators similarly to SineWave.Read
	for i := 0; i < len(p)/4; i++ {
		binary.LittleEndian.PutUint16(p[i*4:], 0)
		binary.LittleEndian.PutUint16(p[i*4+2:], 0)
	}
	return len(p), nil
}

// buildBinauralBeat - builder for the binaural beats module
// GHOST: freq is currently used as baseFreq; beatFreq needs to come from the preset
// When implementing, you'll need to parse/split the frequencies string
// in RunAudio or pass a second freq value from the preset.
func buildBinauralBeat(sampleRate int, freq float64) io.Reader {
	return &BinauralBeat{
		sampleRate: sampleRate,
		baseFreq:   freq, // GHOST: replace with actual base frequency from preset
		beatFreq:   0,    // GHOST: set this from the preset's beat frequency value
		phaseL:     0,
		phaseR:     0,
	}
}

// ========== RunAudio — module dispatcher ==========

// RunAudio plays audio using the wave module matching waveType.
// Added waveType param for module dispatch (was previously single-arg).
func RunAudio(freqStr string, waveType string) {
	freq, err := strconv.ParseFloat(freqStr, 64)
	if err != nil {
		return
	}

	op := &oto.NewContextOptions{}
	op.Format = oto.FormatSignedInt16LE
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.BufferSize = 2048

	ctx, ready, err := oto.NewContext(op)
	if err != nil {
		panic(err)
	}
	<-ready // wait until context is ready

	// Look up the wave module by type, fall back to default
	builder, ok := waveModules[waveType]
	if !ok {
		builder = waveModules["Default (Frequency)"]
	}

	wave := builder(op.SampleRate, freq)

	player := ctx.NewPlayer(wave)
	player.Play()
}
