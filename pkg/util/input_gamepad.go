package util

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
	config [NumOfInputKind]ebiten.GamepadButton
}

func NewGamepadInput() Input {
	return &gamepadInput{
		id: defaultGamepadID,
		config: [NumOfInputKind]ebiten.GamepadButton{
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

func (i *gamepadInput) PressDuration(kind InputKind) int {
	key := i.config[kind]
	return inpututil.GamepadButtonPressDuration(i.id, key)
}

func (i *gamepadInput) JustPressed(kind InputKind) bool {
	key := i.config[kind]
	return inpututil.IsGamepadButtonJustPressed(i.id, key)
}
