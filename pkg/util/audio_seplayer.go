package util

import (
	"bytes"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const sePlayerMaxNum = 100

type sePlayer struct {
	volume  float64
	source  map[int][]byte
	players map[int][sePlayerMaxNum]*audio.Player
}

func NewSEPlayer(sampleRate int, volume float64) SEPlayer {
	return &sePlayer{
		volume:  volume,
		source:  map[int][]byte{},
		players: map[int][sePlayerMaxNum]*audio.Player{},
	}
}

func (se *sePlayer) Load(id int, src []byte) error {
	s, err := mp3.DecodeWithSampleRate(audio.CurrentContext().SampleRate(), bytes.NewReader(src))
	if err != nil {
		return err
	}
	dat, err := io.ReadAll(s)
	if err != nil {
		return err
	}
	se.source[id] = dat
	return nil
}

// Play plays the specified SE.
func (se *sePlayer) Play(id int) {
	src, ok := se.source[id]
	if !ok {
		return
	}
	players := se.players[id]
	for i := range players {
		switch {
		case players[i] == nil:
			p := audio.CurrentContext().NewPlayerFromBytes(src)
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
