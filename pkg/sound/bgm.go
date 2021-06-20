package sound

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/masa213f/stg/resource"
)

type bgmPlayer struct {
	currentBGM resource.BackgroundMusicID
	players    [resource.NumOfBGM]*audio.Player
}

func loadBGM(ctx *audio.Context, resources map[resource.BackgroundMusicID][]byte) (*bgmPlayer, error) {
	bgm := &bgmPlayer{}
	for id, data := range resources {
		s, err := mp3.Decode(ctx, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		l := audio.NewInfiniteLoop(s, s.Length())
		p, err := audio.NewPlayer(ctx, l)
		if err != nil {
			return nil, err
		}
		bgm.players[id] = p
	}
	for i, s := range bgm.players {
		if i != int(resource.BGMNone) && s == nil {
			return nil, fmt.Errorf("BGM[%d] is nil", i)
		}
	}
	return bgm, nil
}

// Reset resets the current BGM and starts the specified BGM from the beginning.
func (bgm *bgmPlayer) Reset(id resource.BackgroundMusicID) {
	if bgm.currentBGM != resource.BGMNone {
		bgm.players[bgm.currentBGM].Pause()
		bgm.players[bgm.currentBGM].Rewind()
	}

	if id != resource.BGMNone {
		bgm.players[id].SetVolume(defaultBGMVolume)
		bgm.players[id].Play()
	}

	bgm.currentBGM = id
}

// Play starts the current BGM if it is paused.
func (bgm *bgmPlayer) Play() {
	if bgm.currentBGM == resource.BGMNone {
		return
	}
	if bgm.players[bgm.currentBGM].IsPlaying() {
		return
	}
	bgm.players[bgm.currentBGM].SetVolume(defaultBGMVolume)
	bgm.players[bgm.currentBGM].Play()
}

// Pause stops the current BGM.
func (bgm *bgmPlayer) Pause() {
	if bgm.currentBGM == resource.BGMNone {
		return
	}
	bgm.players[bgm.currentBGM].Pause()
}
