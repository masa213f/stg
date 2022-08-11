package resource

import (
	_ "embed"
	_ "image/png"
)

const (
	BGMMenu = iota
	BGMPlay
)

// Raw data fo background music.
var (
	//go:embed files/sound/bgm_maoudamashii_fantasy13.mp3
	RawDataBGMMenu []byte
	//go:embed files/sound/bgm_maoudamashii_fantasy15.mp3
	RawDataBGMPlay []byte
)

const (
	SEShot = iota
	SEBomb
	SEHit
	SEDamage
)

// Raw data of sound effects.
var (
	//go:embed files/sound/hitting1.mp3
	RawDataSEShot []byte
	//go:embed files/sound/warp1.mp3
	RawDataSEBomb []byte
	//go:embed files/sound/damage6.mp3
	RawDataSEHit []byte
	//go:embed files/sound/short_bomb.mp3
	RawDataSEDamage []byte
)
