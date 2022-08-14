package util

// MoveDirection represents a direction of movement in a game.
type MoveDirection uint

const (
	MoveNone MoveDirection = iota
	MoveLeft
	MoveRight
	MoveUp
	MoveDown
	MoveUpperLeft
	MoveUpperRight
	MoveLowerLeft
	MoveLowerRight
)

type Control interface {
	Move() MoveDirection
	UpOrDown() MoveDirection
	Select() bool
	Cancel() bool
	Shot() bool
	Bomb() bool
	Pause() bool
}

type InputKind int

const (
	InputKindLeft InputKind = iota
	InputKindRight
	InputKindUp
	InputKindDown
	InputKindOK     // menu
	InputKindCancel // menu
	InputKindShot   // game
	InputKindBomb   // game
	InputKindPause  // game
	NumOfInputKind
)

type Input interface {
	PressDuration(InputKind) int
	JustPressed(InputKind) bool
}

type combinedInput struct {
	rawKeyboardInput Input
	rawGamepadInput  Input
}

func NewCombinedInput(keyboard Input, gamepad Input) Input {
	return &combinedInput{
		rawKeyboardInput: keyboard,
		rawGamepadInput:  gamepad,
	}
}

func (i *combinedInput) PressDuration(kind InputKind) int {
	d1 := i.rawKeyboardInput.PressDuration(kind)
	d2 := i.rawGamepadInput.PressDuration(kind)
	if d1 >= d2 {
		return d1
	}
	return d2
}

func (i *combinedInput) JustPressed(kind InputKind) bool {
	return i.rawKeyboardInput.JustPressed(kind) || i.rawGamepadInput.JustPressed(kind)
}
