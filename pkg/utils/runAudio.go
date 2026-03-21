package utils

import (
	"Frequencer/pkg/models"
	"encoding/binary"
	"io"
	"math"
	"strconv"
	"sync"

	"github.com/ebitengine/oto/v3"
)

var (
	otoCtx       *oto.Context
	otoReady     chan struct{}
	otoOnce      sync.Once
	activePlayer *oto.Player
	playerMu     sync.Mutex
)

func initOtoContext() {
	op := &oto.NewContextOptions{}
	op.Format = oto.FormatSignedInt16LE
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.BufferSize = 2048

	var err error
	otoCtx, otoReady, err = oto.NewContext(op)
	if err != nil {
		panic(err)
	}
}

// --- Wave type module registry ---
// Maps wave type names (from presets) to their builder functions.
// Each builder returns an io.Reader that generates audio samples.
var waveModules = map[string]func(sampleRate int, info models.Presets) io.Reader{
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
func buildSineWave(sampleRate int, info models.Presets) io.Reader {
	freq, err := strconv.ParseFloat(info.Frequencies, 64)
	if err != nil {
		return nil
	}

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

// Read — Generates Left channel at baseFreq, Right channel at baseFreq + beatFreq
func (b *BinauralBeat) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p)/4; i++ {
		// Left channel gets the base frequency
		vL := math.Sin(b.phaseL * 2 * math.Pi)
		b.phaseL += b.baseFreq / float64(b.sampleRate)
		if b.phaseL > 1 {
			b.phaseL -= 1
		}

		// Right channel gets the base + beat frequency
		vR := math.Sin(b.phaseR * 2 * math.Pi)
		b.phaseR += (b.baseFreq + b.beatFreq) / float64(b.sampleRate)
		if b.phaseR > 1 {
			b.phaseR -= 1
		}

		// Scale volume
		sampleL := int16(vL * 0.3 * 32767)
		sampleR := int16(vR * 0.3 * 32767)

		binary.LittleEndian.PutUint16(p[i*4:], uint16(sampleL))
		binary.LittleEndian.PutUint16(p[i*4+2:], uint16(sampleR))
	}
	return len(p), nil
}

// buildBinauralBeat - builder for the binaural beats module
func buildBinauralBeat(sampleRate int, info models.Presets) io.Reader {
	// Try parsing Left frequency or fallback to Frequencies string
	freqStr := info.FrequencieLeft
	if freqStr == "" {
		freqStr = info.Frequencies
	}
	baseFreq, err := strconv.ParseFloat(freqStr, 64)
	if err != nil {
		return nil
	}

	// Try extracting Delta parameter for beat difference
	beatFreq, _ := strconv.ParseFloat(info.Delta, 64)

	return &BinauralBeat{
		sampleRate: sampleRate,
		baseFreq:   baseFreq,
		beatFreq:   beatFreq,
		phaseL:     0,
		phaseR:     0,
	}
}

// ========== RunAudio — module dispatcher ==========

// RunAudio plays audio using the wave module matching waveType.
// Refactored to accept the whole preset struct to dispatch dynamic variables.
func RunAudio(info models.Presets) {
	// Init context only once across the application lifecycle
	otoOnce.Do(initOtoContext)
	<-otoReady // wait until context is ready

	// Look up the wave module by type, fall back to default
	builder, ok := waveModules[info.Type]
	if !ok {
		builder = waveModules["Default (Frequency)"]
	}

	wave := builder(44100, info)
	if wave == nil {
		return // Failed to parse basic freq params in builder
	}

	// Lock to safely handle the active player
	playerMu.Lock()
	defer playerMu.Unlock()

	// If a player is already running, pause it to prevent overlapping frequencies
	if activePlayer != nil {
		activePlayer.Pause()
	}

	activePlayer = otoCtx.NewPlayer(wave)
	activePlayer.Play()
}

// Pause the audio
func PauseAudio() {
	playerMu.Lock()
	defer playerMu.Unlock()

	if activePlayer != nil {
		activePlayer.Pause()
	}
}
