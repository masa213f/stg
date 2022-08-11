package util

type BGMPlayer interface {
	Load(id int, src []byte) error
	Reset(id int)
	Play()
	Pause()
}

type SEPlayer interface {
	Load(id int, src []byte) error
	Play(id int)
}
