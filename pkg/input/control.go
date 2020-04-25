package input

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Default controls
// arrow keys : Move
// z : Shoot
// x : Bomb
// shift : Slowdown movement
// escape : pause

// z : OK
// x : Cancel(Back)

type keyboardConfig struct {
	Left   ebiten.Key
	Right  ebiten.Key
	Up     ebiten.Key
	Down   ebiten.Key
	OK     ebiten.Key // menu
	Cancel ebiten.Key // menu
	Shot   ebiten.Key // game
	Bomb   ebiten.Key // game
	Pause  ebiten.Key // game
}

var defaultKeyboardConfig = keyboardConfig{
	Left:   ebiten.KeyLeft,
	Right:  ebiten.KeyRight,
	Up:     ebiten.KeyUp,
	Down:   ebiten.KeyDown,
	OK:     ebiten.KeyZ,
	Cancel: ebiten.KeyX,
	Shot:   ebiten.KeyZ,
	Bomb:   ebiten.KeyX,
	Pause:  ebiten.KeySpace,
}

type gamepadConfig struct {
	Left   ebiten.GamepadButton
	Right  ebiten.GamepadButton
	Up     ebiten.GamepadButton
	Down   ebiten.GamepadButton
	OK     ebiten.GamepadButton // menu
	Cancel ebiten.GamepadButton // menu
	Shot   ebiten.GamepadButton // game
	Bomb   ebiten.GamepadButton // game
	Pause  ebiten.GamepadButton // game
}

var defaultGamepadConfig = gamepadConfig{
	Left:   ebiten.GamepadButton15,
	Right:  ebiten.GamepadButton13,
	Up:     ebiten.GamepadButton12,
	Down:   ebiten.GamepadButton14,
	OK:     ebiten.GamepadButton0,
	Cancel: ebiten.GamepadButton1,
	Shot:   ebiten.GamepadButton0,
	Bomb:   ebiten.GamepadButton1,
	Pause:  ebiten.GamepadButton2,
}

const defaultGamepadID = 0

type moveControl uint

const (
	bitLeft  moveControl = 1 // 0001
	bitRight moveControl = 2 // 0010
	bitUp    moveControl = 4 // 0100
	bitDown  moveControl = 8 // 1000
)

type MoveAction uint

const (
	MoveNone MoveAction = iota
	MoveLeft
	MoveRight
	MoveUp
	MoveDown
	MoveUpperLeft
	MoveUpperRight
	MoveLowerLeft
	MoveLowerRight
)

// ゲーム中の移動(8方向)
var gameMoveTable = [16]MoveAction{
	MoveNone,       // 0  :0000
	MoveLeft,       // 1  :0001
	MoveRight,      // 2  :0010
	MoveNone,       // 3  :0011
	MoveUp,         // 4  :0100
	MoveUpperLeft,  // 5  :0101
	MoveUpperRight, // 6  :0110
	MoveUp,         // 7  :0111
	MoveDown,       // 8  :1000
	MoveLowerLeft,  // 9  :1001
	MoveLowerRight, // 10 :1010
	MoveDown,       // 11 :1011
	MoveNone,       // 12 :1100
	MoveLeft,       // 13 :1101
	MoveRight,      // 14 :1110
	MoveNone,       // 15 :1111
}

// 8方向
func GameMove() MoveAction {
	var ctrl moveControl

	// keyboard
	if inpututil.KeyPressDuration(defaultKeyboardConfig.Left) > 0 {
		ctrl |= bitLeft
	}
	if inpututil.KeyPressDuration(defaultKeyboardConfig.Right) > 0 {
		ctrl |= bitRight
	}
	if inpututil.KeyPressDuration(defaultKeyboardConfig.Up) > 0 {
		ctrl |= bitUp
	}
	if inpututil.KeyPressDuration(defaultKeyboardConfig.Down) > 0 {
		ctrl |= bitDown
	}

	// gamepad
	if inpututil.GamepadButtonPressDuration(defaultGamepadID, defaultGamepadConfig.Left) > 0 {
		ctrl |= bitLeft
	}
	if inpututil.GamepadButtonPressDuration(defaultGamepadID, defaultGamepadConfig.Right) > 0 {
		ctrl |= bitRight
	}
	if inpututil.GamepadButtonPressDuration(defaultGamepadID, defaultGamepadConfig.Up) > 0 {
		ctrl |= bitUp
	}
	if inpututil.GamepadButtonPressDuration(defaultGamepadID, defaultGamepadConfig.Down) > 0 {
		ctrl |= bitDown
	}

	return gameMoveTable[ctrl]
}

// 上下移動だけ
func MenuUpOrDown() MoveAction {
	up := inpututil.KeyPressDuration(defaultKeyboardConfig.Up)
	down := inpututil.KeyPressDuration(defaultKeyboardConfig.Down)

	if up > 0 && down > 0 {
		return MoveNone
	}
	if up%10 == 1 {
		return MoveUp
	}
	if down%10 == 1 {
		return MoveDown
	}
	return MoveNone
}

func OK() bool {
	return inpututil.IsKeyJustPressed(defaultKeyboardConfig.OK) ||
		inpututil.IsGamepadButtonJustPressed(defaultGamepadID, defaultGamepadConfig.OK)
}

func Cancel() bool {
	return inpututil.IsKeyJustPressed(defaultKeyboardConfig.Cancel) ||
		inpututil.IsGamepadButtonJustPressed(defaultGamepadID, defaultGamepadConfig.Cancel)
}

func Shot() bool {
	return inpututil.KeyPressDuration(defaultKeyboardConfig.Shot) > 0 ||
		inpututil.GamepadButtonPressDuration(defaultGamepadID, defaultGamepadConfig.Shot) > 0
}

func Bomb() bool {
	return inpututil.IsKeyJustPressed(defaultKeyboardConfig.Bomb) ||
		inpututil.IsGamepadButtonJustPressed(defaultGamepadID, defaultGamepadConfig.Bomb)
}

func Pause() bool {
	return inpututil.IsKeyJustPressed(defaultKeyboardConfig.Pause) ||
		inpututil.IsGamepadButtonJustPressed(defaultGamepadID, defaultGamepadConfig.Pause)
}
