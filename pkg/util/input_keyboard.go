//go:build !test

package util

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Default controls
// arrow keys : Move
// z : Shoot
// x : Bomb
// space : pause

// z : OK
// x : Cancel(Back)

const (
	defaultKeyboardConfigLeftKey   = ebiten.KeyLeft
	defaultKeyboardConfigRightKey  = ebiten.KeyRight
	defaultKeyboardConfigUpKey     = ebiten.KeyUp
	defaultKeyboardConfigDownKey   = ebiten.KeyDown
	defaultKeyboardConfigOKKey     = ebiten.KeyZ
	defaultKeyboardConfigCancelKey = ebiten.KeyX
	defaultKeyboardConfigShotKey   = ebiten.KeyZ
	defaultKeyboardConfigBombKey   = ebiten.KeyX
	defaultKeyboardConfigPauseKey  = ebiten.KeySpace
)

type keyboardInput struct {
	config [NumOfInputKind]ebiten.Key
}

func NewKeyboardInput() Input {
	return &keyboardInput{
		config: [NumOfInputKind]ebiten.Key{
			defaultKeyboardConfigLeftKey,
			defaultKeyboardConfigRightKey,
			defaultKeyboardConfigUpKey,
			defaultKeyboardConfigDownKey,
			defaultKeyboardConfigOKKey,
			defaultKeyboardConfigCancelKey,
			defaultKeyboardConfigShotKey,
			defaultKeyboardConfigBombKey,
			defaultKeyboardConfigPauseKey,
		},
	}
}

func (i *keyboardInput) PressDuration(kind InputKind) int {
	key := i.config[kind]
	return inpututil.KeyPressDuration(key)
}

func (i *keyboardInput) JustPressed(kind InputKind) bool {
	key := i.config[kind]
	return inpututil.IsKeyJustPressed(key)
}
