package util

type control struct {
	input Input
}

func NewControl(input Input) Control {
	return &control{
		input: input,
	}
}

const (
	inputLeft  = 1 // 0001
	inputRight = 2 // 0010
	inputUp    = 4 // 0100
	inputDown  = 8 // 1000
)

// table of input values and directions of movement in the game.
var gameMoveTable = [16]MoveDirection{
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

func (c *control) Move() MoveDirection {
	var in uint
	if c.input.PressDuration(InputKindLeft) > 0 {
		in |= inputLeft
	}
	if c.input.PressDuration(InputKindRight) > 0 {
		in |= inputRight
	}
	if c.input.PressDuration(InputKindUp) > 0 {
		in |= inputUp
	}
	if c.input.PressDuration(InputKindDown) > 0 {
		in |= inputDown
	}
	return gameMoveTable[in]
}

// UpOrDown returns the vertical movement.
func (c *control) UpOrDown() MoveDirection {
	up := c.input.PressDuration(InputKindUp)
	down := c.input.PressDuration(InputKindDown)

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

func (c *control) Select() bool {
	return c.input.JustPressed(InputKindOK)
}

func (c *control) Cancel() bool {
	return c.input.JustPressed(InputKindCancel)
}

func (c *control) Shot() bool {
	return c.input.PressDuration(InputKindOK) > 0
}

func (c *control) Bomb() bool {
	return c.input.PressDuration(InputKindCancel) > 0
}

func (c *control) Pause() bool {
	return c.input.JustPressed(InputKindPause)
}
