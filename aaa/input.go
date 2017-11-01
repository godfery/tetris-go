package aaa

import (
	"github.com/hajimehoshi/ebiten"
)

// Dir represents a direction.
type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

type touchState int

const (
	touchStateNone touchState = iota
	touchStatePressing
	touchStateSettled
	touchStateInvalid
	n
)

// String returns a string representing the direction.
func (d Dir) String() string {
	switch d {
	case DirUp:
		return "Up"
	case DirRight:
		return "Right"
	case DirDown:
		return "Down"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

// Vector returns a [-1, 1] value for each axis.
func (d Dir) Vector() (x, y int) {
	switch d {
	case DirUp:
		return 0, -1
	case DirRight:
		return 1, 0
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	}
	panic("not reach")
}

// Input represents the current key states.
type Input struct {
	keyState map[ebiten.Key]int

	mouseState    mouseState
	mouseInitPosX int
	mouseInitPosY int
	mouseDir      Dir

	touchState    touchState
	touchID       int
	touchInitPosX int
	touchInitPosY int
	touchLastPosX int
	touchLastPosY int
	touchDir      Dir
}

// NewInput generates a new Input object.
func NewInput() *Input {
	return &Input{
		keyState: map[ebiten.Key]int{},
	}
}

var (
	dirKeys = map[ebiten.Key]Dir{
		ebiten.KeyUp:    DirUp,
		ebiten.KeyRight: DirRight,
		ebiten.KeyDown:  DirDown,
		ebiten.KeyLeft:  DirLeft,
	}
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func vecToDir(dx, dy int) (Dir, bool) {
	if abs(dx) < 4 && abs(dy) < 4 {
		return 0, false
	}
	if abs(dx) < abs(dy) {
		if dy < 0 {
			return DirUp, true
		}
		return DirDown, true
	} else {
		if dx < 0 {
			return DirLeft, true
		}
		return DirRight, true
	}
}

// Update updates the current input states.
func (i *Input) Update() {
	for k := range dirKeys {
		if ebiten.IsKeyPressed(k) {
			i.keyState[k]++
		} else {
			i.keyState[k] = 0
		}
	}
	switch i.mouseState {
	case mouseStateNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			i.mouseInitPosX = x
			i.mouseInitPosY = y
			i.mouseState = mouseStatePressing
		}
	case mouseStatePressing:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			dx := x - i.mouseInitPosX
			dy := y - i.mouseInitPosY
			d, ok := vecToDir(dx, dy)
			if !ok {
				i.mouseState = mouseStateNone
				break
			}
			i.mouseDir = d
			i.mouseState = mouseStateSettled
		}
	case mouseStateSettled:
		i.mouseState = mouseStateNone
	}
	switch i.touchState {
	case touchStateNone:
		ts := ebiten.Touches()
		if len(ts) == 1 {
			i.touchID = ts[0].ID()
			x, y := ts[0].Position()
			i.touchInitPosX = x
			i.touchInitPosY = y
			i.touchLastPosX = x
			i.touchLastPosX = y
			i.touchState = touchStatePressing
		}
	case touchStatePressing:
		ts := ebiten.Touches()
		if len(ts) >= 2 {
			break
		}
		if len(ts) == 1 {
			if ts[0].ID() != i.touchID {
				i.touchState = touchStateInvalid
			} else {
				x, y := ts[0].Position()
				i.touchLastPosX = x
				i.touchLastPosY = y
			}
			break
		}
		if len(ts) == 0 {
			dx := i.touchLastPosX - i.touchInitPosX
			dy := i.touchLastPosY - i.touchInitPosY
			d, ok := vecToDir(dx, dy)
			if !ok {
				i.touchState = touchStateNone
				break
			}
			i.touchDir = d
			i.touchState = touchStateSettled
		}
	case touchStateSettled:
		i.touchState = touchStateNone
	case touchStateInvalid:
		if len(ebiten.Touches()) == 0 {
			i.touchState = touchStateNone
		}
	}
}

// Dir returns a currently pressed direction.
// Dir returns false if no direction key is pressed.
func (i *Input) Dir() (Dir, bool) {
	for k, d := range dirKeys {
		if i.keyState[k] == 1 {
			return d, true
		}
	}
	if i.mouseState == mouseStateSettled {
		return i.mouseDir, true
	}
	if i.touchState == touchStateSettled {
		return i.touchDir, true
	}
	return 0, false
}
