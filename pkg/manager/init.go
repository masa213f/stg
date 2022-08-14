package manager

import (
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

const (
	sampleRate       = 44100
	defaultBGMVolume = 0.1
	defaultSEVolume  = 0.1
)

var (
	audio util.AudioPlayer
)

func init() {
	util.InitAudio(sampleRate)
	audio = util.NewAudioPlayer(defaultBGMVolume, defaultSEVolume)

	seList := map[int][]byte{
		resource.SEShot:   resource.RawDataSEShot,
		resource.SEBomb:   resource.RawDataSEBomb,
		resource.SEHit:    resource.RawDataSEHit,
		resource.SEDamage: resource.RawDataSEDamage,
	}
	for id, src := range seList {
		if err := audio.LoadSE(id, src); err != nil {
			panic(err)
		}
	}

	bgmList := map[int][]byte{
		resource.BGMMenu: resource.RawDataBGMMenu,
		resource.BGMPlay: resource.RawDataBGMPlay,
	}
	for id, src := range bgmList {
		if err := audio.LoadBGM(id, src); err != nil {
			panic(err)
		}
	}
}
