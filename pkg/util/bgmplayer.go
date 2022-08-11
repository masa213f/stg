package util

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

func InitAudio(sampleRate int) {
	audio.NewContext(sampleRate)
}

type bgmPlayer struct {
	volume     float64
	currentBGM int
	players    map[int]*audio.Player
}

func NewBGMPlayer(sampleRate int, volume float64) BGMPlayer {
	return &bgmPlayer{
		volume:  volume,
		players: map[int]*audio.Player{},
	}
}

func (bgm *bgmPlayer) Load(id int, src []byte) error {
	dat, err := mp3.DecodeWithSampleRate(audio.CurrentContext().SampleRate(), bytes.NewReader(src))
	if err != nil {
		return err
	}
	p, err := audio.CurrentContext().NewPlayer(audio.NewInfiniteLoop(dat, dat.Length()))
	if err != nil {
		return err
	}
	bgm.players[id] = p
	return nil
}

// Reset resets the current BGM and starts the specified BGM from the beginning.
func (bgm *bgmPlayer) Reset(id int) {
	if p, ok := bgm.players[bgm.currentBGM]; ok {
		p.Pause()
		p.Rewind()
	}
	if p, ok := bgm.players[id]; ok {
		p.SetVolume(bgm.volume)
		p.Play()
	}
	bgm.currentBGM = id
}

// Play starts the current BGM if it is paused.
func (bgm *bgmPlayer) Play() {
	if p, ok := bgm.players[bgm.currentBGM]; ok {
		if p.IsPlaying() {
			return
		}
		p.SetVolume(bgm.volume)
		p.Play()
	}
}

// Pause stops the current BGM.
func (bgm *bgmPlayer) Pause() {
	if p, ok := bgm.players[bgm.currentBGM]; ok {
		p.Pause()
	}
}
