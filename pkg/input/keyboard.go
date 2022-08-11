//go:build !test

package input

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
	config [numOfInputKind]ebiten.Key
}

func init() {
	setKeyboardInput(newKeyboardInput())
}

func newKeyboardInput() rawInput {
	return &keyboardInput{
		config: [numOfInputKind]ebiten.Key{
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

func (i *keyboardInput) PressDuration(kind inputKind) int {
	key := i.config[kind]
	return inpututil.KeyPressDuration(key)
}

func (i *keyboardInput) JustPressed(kind inputKind) bool {
	key := i.config[kind]
	return inpututil.IsKeyJustPressed(key)
}
