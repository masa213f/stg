package resource

import (
	_ "embed"
	"fmt"
	_ "image/png"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

// Raw data fo background music.
var (
	//go:embed files/sound/bgm_maoudamashii_fantasy13.mp3
	rawDataBGMMenu []byte
	//go:embed files/sound/bgm_maoudamashii_fantasy15.mp3
	rawDataBGMPlay []byte
)

// Raw data of sound effects.
var (
	//go:embed files/sound/hitting1.mp3
	rawDataSEShot []byte
	//go:embed files/sound/warp1.mp3
	rawDataSEBomb []byte
	//go:embed files/sound/damage6.mp3
	rawDataSEHit []byte
	//go:embed files/sound/short_bomb.mp3
	rawDataSEDamage []byte
)

const (
	BGMVolume  = 1
	SEVolume   = 1
	sampleRate = 44100
)

type BackgroundMusicID int

const (
	BGMNone BackgroundMusicID = iota
	BGMMenu
	BGMPlay
	NumOfBGM
)

type SoundEffectID int

const (
	SEShot SoundEffectID = iota
	SEBomb
	SEHit
	SEDamage
	NumOfSE
)

var (
	audioContext *audio.Context
	BGM          *backgroundMusic
	SE           *soundEffect
)

type soundEffect struct {
	source [NumOfSE][]byte
}

func loadSoundEffect(resources map[SoundEffectID][]byte) (*soundEffect, error) {
	se := &soundEffect{}
	for k, v := range resources {
		s, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(v))
		if err != nil {
			return nil, err
		}
		src, err := ioutil.ReadAll(s)
		if err != nil {
			return nil, err
		}
		se.source[k] = src
	}
	for i, s := range se.source {
		if s == nil {
			return nil, fmt.Errorf("SE[%d] is nil", i)
		}
	}
	return se, nil
}

// Play plays the specified SE.
func (s *soundEffect) Play(id SoundEffectID) {
	p, _ := audio.NewPlayerFromBytes(audioContext, s.source[id])
	p.SetVolume(SEVolume)
	p.Play()
	// TODO: Should the player be closed?
}

type backgroundMusic struct {
	currentBGM BackgroundMusicID
	players    [NumOfBGM]*audio.Player
}

func loadBackgroundMusic(resources map[BackgroundMusicID][]byte) (*backgroundMusic, error) {
	bgm := &backgroundMusic{}
	for k, v := range resources {
		s, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(v))
		if err != nil {
			return nil, err
		}
		l := audio.NewInfiniteLoop(s, s.Length())
		p, err := audio.NewPlayer(audioContext, l)
		if err != nil {
			return nil, err
		}
		bgm.players[k] = p
	}
	for i, s := range bgm.players {
		if i != int(BGMNone) && s == nil {
			return nil, fmt.Errorf("BGM[%d] is nil", i)
		}
	}
	return bgm, nil
}

// Reset resets the current BGM and starts the specified BGM from the beginning.
func (s *backgroundMusic) Reset(id BackgroundMusicID) {
	if s.currentBGM != BGMNone {
		s.players[s.currentBGM].Pause()
		s.players[s.currentBGM].Rewind()
	}

	if id != BGMNone {
		s.players[id].Play()
	}

	s.currentBGM = id
}

// Play starts the current BGM if it is paused.
func (s *backgroundMusic) Play() {
	if s.currentBGM == BGMNone {
		return
	}
	if s.players[s.currentBGM].IsPlaying() {
		return
	}
	s.players[s.currentBGM].Play()
}

// Pause stops the current BGM.
func (s *backgroundMusic) Pause() {
	if s.currentBGM == BGMNone {
		return
	}
	s.players[s.currentBGM].Pause()
}

func init() {
	audioContext, _ = audio.NewContext(sampleRate)

	var err error
	SE, err = loadSoundEffect(map[SoundEffectID][]byte{
		SEShot:   rawDataSEShot,
		SEBomb:   rawDataSEBomb,
		SEHit:    rawDataSEHit,
		SEDamage: rawDataSEDamage,
	})
	if err != nil {
		panic(err)
	}

	BGM, err = loadBackgroundMusic(map[BackgroundMusicID][]byte{
		BGMMenu: rawDataBGMMenu,
		BGMPlay: rawDataBGMPlay,
	})
	if err != nil {
		panic(err)
	}
}
