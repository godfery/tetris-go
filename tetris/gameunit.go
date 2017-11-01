package tetris

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	BackgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	FrameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

type GameUnit struct {
	input *Input
	// board      *Board
	boardImage *ebiten.Image
	tObject    *TetrisObject
}

const (
	ScreenWidth  = 410
	ScreenHeight = 300
	boardSize    = 5
	size         = 10
)

const (
	tileSize   = 20
	tileMargin = 2
)

// NewGame generates a new Game object.
func NewGame() (*GameUnit, error) {
	g := &GameUnit{
		input: NewInput(),
	}
	var err error
	// g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (gggg *GameUnit) Draw() {

	// fmt.Println("draw")
	gggg.tObject.MoveByDir(gggg.input)
	gggg.boardImage.Fill(FrameColor)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			v := 0
			op := &ebiten.DrawImageOptions{}
			x := i*tileSize + (i+1)*tileMargin
			y := j*tileSize + (j+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			if gggg.tObject.LineCage[j][i] > 0 {
				v = 4096
			}
			r, g, b, a := colorToScale(tileBackgroundColor(v))
			op.ColorM.Scale(r, g, b, a)
			tileImage, _ := ebiten.NewImage(tileSize, tileSize, ebiten.FilterNearest)
			tileImage.Fill(color.White)
			gggg.boardImage.DrawImage(tileImage, op)
		}
	}
}
func frameSize() (int, int) {
	x := size*tileSize + (size+1)*tileMargin
	y := x
	return x, y
}

func (gggg *GameUnit) Update(screen *ebiten.Image) {
	gggg.input.Update()
	// fmt.Println(gggg.input.Dir())
	if gggg.boardImage == nil {
		w, h := frameSize()
		gggg.boardImage, _ = ebiten.NewImage(w, h, ebiten.FilterNearest)
	}
	if gggg.tObject == nil {
		gggg.tObject = NewTetris()
	}

	screen.Fill(BackgroundColor)

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := gggg.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	// fmt.Println(x, y, sw, sh, bw, bh)
	op.GeoM.Translate(float64(x), float64(y))

	gggg.Draw()
	//
	gggg.tObject.Draw(gggg.boardImage)
	screen.DrawImage(gggg.boardImage, op)
}
