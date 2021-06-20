package sound

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/masa213f/stg/resource"
)

const sePlayerMaxNum = 100

type sePlayer struct {
	volume  float64
	source  [resource.NumOfSE][]byte
	players [resource.NumOfSE][sePlayerMaxNum]*audio.Player
}

func loadSE(ctx *audio.Context, resources map[resource.SoundEffectID][]byte) (*sePlayer, error) {
	se := &sePlayer{}
	for id, data := range resources {
		s, err := mp3.Decode(ctx, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		src, err := ioutil.ReadAll(s)
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
			p := audio.NewPlayerFromBytes(audioContext, se.source[id])
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
