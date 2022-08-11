package input

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "control")
}

type mockRawInput struct {
	data map[inputKind]int
}

func newMockRawInput(data map[inputKind]int) rawInput {
	if data == nil {
		data = map[inputKind]int{}
	}
	return &mockRawInput{
		data: data,
	}
}

func (i *mockRawInput) PressDuration(kind inputKind) int {
	return i.data[kind]
}

func (i *mockRawInput) JustPressed(kind inputKind) bool {
	return i.data[kind] == 1
}

var _ = Describe("control", func() {
	It("should return the proper direction", func() {
		testcase := []*struct {
			keyboard map[inputKind]int
			gamepad  map[inputKind]int
			expect   MoveAction
		}{
			{
				keyboard: map[inputKind]int{},
				gamepad:  map[inputKind]int{},
				expect:   MoveNone,
			},

			// keyboard
			{
				keyboard: map[inputKind]int{inputKindLeft: 1},
				expect:   MoveLeft,
			},
			{
				keyboard: map[inputKind]int{inputKindRight: 2},
				expect:   MoveRight,
			},
			{
				keyboard: map[inputKind]int{inputKindLeft: 1, inputKindRight: 2},
				expect:   MoveNone,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3},
				expect:   MoveUp,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindLeft: 1},
				expect:   MoveUpperLeft,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindRight: 2},
				expect:   MoveUpperRight,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindLeft: 1, inputKindRight: 2},
				expect:   MoveUp,
			},
			{
				keyboard: map[inputKind]int{inputKindDown: 3},
				expect:   MoveDown,
			},
			{
				keyboard: map[inputKind]int{inputKindDown: 3, inputKindLeft: 1},
				expect:   MoveLowerLeft,
			},
			{
				keyboard: map[inputKind]int{inputKindDown: 3, inputKindRight: 2},
				expect:   MoveLowerRight,
			},
			{
				keyboard: map[inputKind]int{inputKindDown: 3, inputKindLeft: 1, inputKindRight: 2},
				expect:   MoveDown,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindDown: 3},
				expect:   MoveNone,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindDown: 3, inputKindLeft: 1},
				expect:   MoveLeft,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindDown: 3, inputKindRight: 2},
				expect:   MoveRight,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindDown: 3, inputKindLeft: 1, inputKindRight: 2},
				expect:   MoveNone,
			},

			// gamepad
			{
				gamepad: map[inputKind]int{inputKindLeft: 1},
				expect:  MoveLeft,
			},
			{
				gamepad: map[inputKind]int{inputKindRight: 2},
				expect:  MoveRight,
			},
			{
				gamepad: map[inputKind]int{inputKindLeft: 1, inputKindRight: 2},
				expect:  MoveNone,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3},
				expect:  MoveUp,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3, inputKindLeft: 1},
				expect:  MoveUpperLeft,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3, inputKindRight: 2},
				expect:  MoveUpperRight,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3, inputKindLeft: 1, inputKindRight: 2},
				expect:  MoveUp,
			},
			{
				gamepad: map[inputKind]int{inputKindDown: 3},
				expect:  MoveDown,
			},
			{
				gamepad: map[inputKind]int{inputKindDown: 3, inputKindLeft: 1},
				expect:  MoveLowerLeft,
			},
			{
				gamepad: map[inputKind]int{inputKindDown: 3, inputKindRight: 2},
				expect:  MoveLowerRight,
			},
			{
				gamepad: map[inputKind]int{inputKindDown: 3, inputKindLeft: 1, inputKindRight: 2},
				expect:  MoveDown,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3, inputKindDown: 3},
				expect:  MoveNone,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3, inputKindDown: 3, inputKindLeft: 1},
				expect:  MoveLeft,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3, inputKindDown: 3, inputKindRight: 2},
				expect:  MoveRight,
			},
			{
				gamepad: map[inputKind]int{inputKindUp: 3, inputKindDown: 3, inputKindLeft: 1, inputKindRight: 2},
				expect:  MoveNone,
			},

			// keyboard + gamepad
			{
				keyboard: map[inputKind]int{inputKindUp: 3},
				gamepad:  map[inputKind]int{inputKindLeft: 1},
				expect:   MoveUpperLeft,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindRight: 2},
				gamepad:  map[inputKind]int{inputKindDown: 3},
				expect:   MoveRight,
			},
			{
				keyboard: map[inputKind]int{inputKindUp: 3, inputKindLeft: 1},
				gamepad:  map[inputKind]int{inputKindDown: 3, inputKindRight: 2},
				expect:   MoveNone,
			},
		}

		for i, tt := range testcase {
			setKeyboardInput(newMockRawInput(tt.keyboard))
			setGamepadInput(newMockRawInput(tt.gamepad))
			By(fmt.Sprintf("%d", i))
			Expect(Move()).To(Equal(tt.expect))
		}
	})

	// TODO: Add test for OK(), Cancel(), and so on.
})
