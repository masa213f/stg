package sound

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/masa213f/stg/resource"
)

const (
	defaultBGMVolume = 0.1
	defaultSEVolume  = 0.1
	sampleRate       = 44100
)

var (
	audioContext *audio.Context
	BGM          *bgmPlayer
	SE           *sePlayer
)

func init() {
	audioContext = audio.NewContext(sampleRate)

	var err error
	SE, err = loadSE(audioContext, map[resource.SoundEffectID][]byte{
		resource.SEShot:   resource.RawDataSEShot,
		resource.SEBomb:   resource.RawDataSEBomb,
		resource.SEHit:    resource.RawDataSEHit,
		resource.SEDamage: resource.RawDataSEDamage,
	})
	if err != nil {
		panic(err)
	}
	SE.SetVolume(defaultSEVolume)

	BGM, err = loadBGM(audioContext, map[resource.BackgroundMusicID][]byte{
		resource.BGMMenu: resource.RawDataBGMMenu,
		resource.BGMPlay: resource.RawDataBGMPlay,
	})
	if err != nil {
		panic(err)
	}
}
