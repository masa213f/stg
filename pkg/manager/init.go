package manager

import (
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

const (
	defaultBGMVolume = 0.1
	defaultSEVolume  = 0.1
	sampleRate       = 44100
)

var (
	BGM util.BGMPlayer
	SE  util.SEPlayer
)

func init() {
	util.InitAudio(sampleRate)

	seList := map[int][]byte{
		resource.SEShot:   resource.RawDataSEShot,
		resource.SEBomb:   resource.RawDataSEBomb,
		resource.SEHit:    resource.RawDataSEHit,
		resource.SEDamage: resource.RawDataSEDamage,
	}
	SE = util.NewSEPlayer(sampleRate, defaultSEVolume)
	for id, src := range seList {
		err := SE.Load(id, src)
		if err != nil {
			panic(err)
		}
	}

	bgmList := map[int][]byte{
		resource.BGMMenu: resource.RawDataBGMMenu,
		resource.BGMPlay: resource.RawDataBGMPlay,
	}
	BGM = util.NewBGMPlayer(sampleRate, defaultBGMVolume)
	for id, src := range bgmList {
		err := BGM.Load(id, src)
		if err != nil {
			panic(err)
		}
	}
}
