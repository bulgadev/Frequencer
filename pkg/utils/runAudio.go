package utils

import (
	"encoding/binary"
	"math"
	"strconv"

	"github.com/ebitengine/oto/v3"
)

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

func buildAudio(sampleRate int, freq float64) *SineWave {
	return &SineWave{
		sampleRate: sampleRate,
		freq:       freq,
		phase:      0,
	}
}

func RunAudio(freqStr string) {
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

	wave := buildAudio(op.SampleRate, freq)

	player := ctx.NewPlayer(wave)
	player.Play()
}
