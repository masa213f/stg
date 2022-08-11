package input

type inputKind int

const (
	inputKindLeft inputKind = iota
	inputKindRight
	inputKindUp
	inputKindDown
	inputKindOK     // menu
	inputKindCancel // menu
	inputKindShot   // game
	inputKindBomb   // game
	inputKindPause  // game
	numOfInputKind
)

type rawInput interface {
	PressDuration(inputKind) int
	JustPressed(inputKind) bool
}

var (
	rawKeyboardInput rawInput
	rawGamepadInput  rawInput
)

func setKeyboardInput(raw rawInput) {
	rawKeyboardInput = raw
}

func setGamepadInput(raw rawInput) {
	rawGamepadInput = raw
}

// moveControl represents a direction of input.
type moveControl uint

const (
	bitLeft  moveControl = 1 // 0001
	bitRight moveControl = 2 // 0010
	bitUp    moveControl = 4 // 0100
	bitDown  moveControl = 8 // 1000
)

// MoveAction represents a direction of movement in a game.
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

// table of input values and directions of movement in the game.
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

// Move returns movement in 8 directions.
func Move() MoveAction {
	var ctrl moveControl

	if rawKeyboardInput.PressDuration(inputKindLeft) > 0 ||
		rawGamepadInput.PressDuration(inputKindLeft) > 0 {
		ctrl |= bitLeft
	}
	if rawKeyboardInput.PressDuration(inputKindRight) > 0 ||
		rawGamepadInput.PressDuration(inputKindRight) > 0 {
		ctrl |= bitRight
	}
	if rawKeyboardInput.PressDuration(inputKindUp) > 0 ||
		rawGamepadInput.PressDuration(inputKindUp) > 0 {
		ctrl |= bitUp
	}
	if rawKeyboardInput.PressDuration(inputKindDown) > 0 ||
		rawGamepadInput.PressDuration(inputKindDown) > 0 {
		ctrl |= bitDown
	}

	return gameMoveTable[ctrl]
}

// UpOrDown returns the vertical movement.
func UpOrDown() MoveAction {
	// TODO: handle the gamepad input.
	up := rawKeyboardInput.PressDuration(inputKindUp)
	down := rawKeyboardInput.PressDuration(inputKindDown)

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
	return rawKeyboardInput.JustPressed(inputKindOK) ||
		rawGamepadInput.JustPressed(inputKindOK)
}

func Cancel() bool {
	return rawKeyboardInput.JustPressed(inputKindCancel) ||
		rawGamepadInput.JustPressed(inputKindCancel)
}

func Shot() bool {
	return rawKeyboardInput.PressDuration(inputKindShot) > 0 ||
		rawGamepadInput.PressDuration(inputKindShot) > 0
}

func Bomb() bool {
	return rawKeyboardInput.PressDuration(inputKindBomb) > 0 ||
		rawGamepadInput.PressDuration(inputKindBomb) > 0
}

func Pause() bool {
	return rawKeyboardInput.JustPressed(inputKindPause) ||
		rawGamepadInput.JustPressed(inputKindPause)
}
