package util

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestControl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "control")
}

type mockInput struct {
	data map[InputKind]int
}

func (i *mockInput) PressDuration(kind InputKind) int {
	return i.data[kind]
}

func (i *mockInput) JustPressed(kind InputKind) bool {
	return i.data[kind] == 1
}

var _ = Describe("control", func() {
	It("should return the proper direction", func() {
		testcase := []*struct {
			input  map[InputKind]int
			expect MoveDirection
		}{
			{
				input:  map[InputKind]int{},
				expect: MoveNone,
			},
			{
				input:  map[InputKind]int{InputKindLeft: 1},
				expect: MoveLeft,
			},
			{
				input:  map[InputKind]int{InputKindRight: 2},
				expect: MoveRight,
			},
			{
				input:  map[InputKind]int{InputKindLeft: 1, InputKindRight: 2},
				expect: MoveNone,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3},
				expect: MoveUp,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3, InputKindLeft: 1},
				expect: MoveUpperLeft,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3, InputKindRight: 2},
				expect: MoveUpperRight,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3, InputKindLeft: 1, InputKindRight: 2},
				expect: MoveUp,
			},
			{
				input:  map[InputKind]int{InputKindDown: 3},
				expect: MoveDown,
			},
			{
				input:  map[InputKind]int{InputKindDown: 3, InputKindLeft: 1},
				expect: MoveLowerLeft,
			},
			{
				input:  map[InputKind]int{InputKindDown: 3, InputKindRight: 2},
				expect: MoveLowerRight,
			},
			{
				input:  map[InputKind]int{InputKindDown: 3, InputKindLeft: 1, InputKindRight: 2},
				expect: MoveDown,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3, InputKindDown: 3},
				expect: MoveNone,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3, InputKindDown: 3, InputKindLeft: 1},
				expect: MoveLeft,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3, InputKindDown: 3, InputKindRight: 2},
				expect: MoveRight,
			},
			{
				input:  map[InputKind]int{InputKindUp: 3, InputKindDown: 3, InputKindLeft: 1, InputKindRight: 2},
				expect: MoveNone,
			},
		}

		mock := &mockInput{}
		ctrl := NewControl(mock)
		for i, tt := range testcase {
			By(fmt.Sprintf("%d", i))
			mock.data = tt.input
			Expect(ctrl.Move()).To(Equal(tt.expect))
		}
	})

	// TODO: Add test for OK(), Cancel(), and so on.
})
