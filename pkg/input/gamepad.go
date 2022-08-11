//go:build !test

package input

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const defaultGamepadID = 0

const (
	defaultGamepadConfigLeft   = ebiten.GamepadButton15
	defaultGamepadConfigRight  = ebiten.GamepadButton13
	defaultGamepadConfigUp     = ebiten.GamepadButton12
	defaultGamepadConfigDown   = ebiten.GamepadButton14
	defaultGamepadConfigOK     = ebiten.GamepadButton0
	defaultGamepadConfigCancel = ebiten.GamepadButton1
	defaultGamepadConfigShot   = ebiten.GamepadButton0
	defaultGamepadConfigBomb   = ebiten.GamepadButton1
	defaultGamepadConfigPause  = ebiten.GamepadButton2
)

type gamepadInput struct {
	id     ebiten.GamepadID
	config [numOfInputKind]ebiten.GamepadButton
}

func init() {
	setGamepadInput(newGamepadInput())
}

func newGamepadInput() rawInput {
	return &gamepadInput{
		id: defaultGamepadID,
		config: [numOfInputKind]ebiten.GamepadButton{
			defaultGamepadConfigLeft,
			defaultGamepadConfigRight,
			defaultGamepadConfigUp,
			defaultGamepadConfigDown,
			defaultGamepadConfigOK,
			defaultGamepadConfigCancel,
			defaultGamepadConfigShot,
			defaultGamepadConfigBomb,
			defaultGamepadConfigPause,
		},
	}
}

func (i *gamepadInput) PressDuration(kind inputKind) int {
	key := i.config[kind]
	return inpututil.GamepadButtonPressDuration(i.id, key)
}

func (i *gamepadInput) JustPressed(kind inputKind) bool {
	key := i.config[kind]
	return inpututil.IsGamepadButtonJustPressed(i.id, key)
}
