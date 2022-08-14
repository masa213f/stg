package util

import (
	"bytes"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type AudioPlayer interface {
	LoadBGM(bgmId int, src []byte) error
	ResetBGM(bgmId int)
	PlayBGM()
	PauseBGM()

	LoadSE(seId int, src []byte) error
	PlaySE(seId int)
}

type audioPlayer struct {
	bgm *bgmPlayer
	se  *sePlayer
}

type bgmPlayer struct {
	volume     float64
	currentBGM int
	players    map[int]*audio.Player
}

const sePlayerMaxNum = 100

type sePlayer struct {
	volume  float64
	source  map[int][]byte
	players map[int][sePlayerMaxNum]*audio.Player
}

func InitAudio(sampleRate int) {
	audio.NewContext(sampleRate)
}

func NewAudioPlayer(bgmVolume, seVolume float64) AudioPlayer {
	return &audioPlayer{
		bgm: newBGMPlayer(bgmVolume),
		se:  newSEPlayer(seVolume),
	}
}

func (p *audioPlayer) LoadBGM(bgmId int, src []byte) error {
	return p.bgm.Load(bgmId, src)
}

func (p *audioPlayer) ResetBGM(bgmId int) {
	p.bgm.Reset(bgmId)
}

func (p *audioPlayer) PlayBGM() {
	p.bgm.Play()
}

func (p *audioPlayer) PauseBGM() {
	p.bgm.Pause()
}

func (p *audioPlayer) LoadSE(seId int, src []byte) error {
	return p.se.Load(seId, src)
}

func (p *audioPlayer) PlaySE(seId int) {
	p.se.Play(seId)
}

func newBGMPlayer(volume float64) *bgmPlayer {
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

func newSEPlayer(volume float64) *sePlayer {
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
