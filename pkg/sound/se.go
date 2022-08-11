package sound

import (
	"bytes"
	"fmt"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/masa213f/stg/resource"
)

const sePlayerMaxNum = 100

type sePlayer struct {
	audioContext *audio.Context
	volume       float64
	source       [resource.NumOfSE][]byte
	players      [resource.NumOfSE][sePlayerMaxNum]*audio.Player
}

func loadSE(audioContext *audio.Context, sampleRate int, resources map[resource.SoundEffectID][]byte) (*sePlayer, error) {
	se := &sePlayer{
		audioContext: audioContext,
	}
	for id, data := range resources {
		s, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		src, err := io.ReadAll(s)
		if err != nil {
			return nil, err
		}
		se.source[id] = src
	}
	for i, s := range se.source {
		if s == nil {
			return nil, fmt.Errorf("SE[%d] is nil", i)
		}
	}
	return se, nil
}

func (se *sePlayer) SetVolume(volume float64) {
	se.volume = volume
}

// Play plays the specified SE.
func (se *sePlayer) Play(id resource.SoundEffectID) {
	players := &se.players[id]
	for i := range players {
		switch {
		case players[i] == nil:
			p := se.audioContext.NewPlayerFromBytes(se.source[id])
			p.SetVolume(se.volume)
			p.Play()
			players[i] = p
			return
		case !players[i].IsPlaying():
			p := players[i]
			p.Rewind()
			p.SetVolume(se.volume)
			p.Play()
			return
		}
	}
}
