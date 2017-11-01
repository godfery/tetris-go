package tetris

import (
	"math/rand"
	"time"
)

func Generate_Randnum(param int) int {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := rnd.Intn(param)

	return vcode
}

func InitCage() areas {
	var allAreas areas
	var area1 area
	var b [4][4]int
	b = [4][4]int{
		{1, 1, 1, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	area1.param = b
	area1.xLen = 3
	area1.yLen = 2
	allAreas.all[0] = area1
	allAreas.rotation[0][0] = area1

	b = [4][4]int{
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0}}
	area1.param = b
	area1.xLen = 2
	area1.yLen = 3

	allAreas.rotation[0][1] = area1

	b = [4][4]int{
		{0, 0, 1, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	area1.param = b
	area1.xLen = 3
	area1.yLen = 2

	allAreas.rotation[0][2] = area1

	b = [4][4]int{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0}}
	area1.param = b
	area1.xLen = 3
	area1.yLen = 2

	allAreas.rotation[0][3] = area1

	var a [4][4]int
	a[1][1] = 1
	a[1][2] = 1
	a[2][2] = 1
	a[2][1] = 1
	area1.param = a
	area1.xLen = 2
	area1.yLen = 2
	allAreas.all[1] = area1

	return allAreas
}
func checkCanMove(i int) bool {
	if i < 0 || i > 9 {
		return false
	}
	return true
}

// func (to *TetrisObject) CageDraw(image *ebiten.Image) {
//
// 	// fmt.Println(to.LineCage)
// 	for xx := 0; xx < size; xx++ {
// 		for jj := 0; jj < size; jj++ {
// 			area := to.LineCage
//
// 			if area[xx][jj] > 0 {
// 				j := xx
// 				i := jj
// 				v := 4096
//
// 				op := &ebiten.DrawImageOptions{}
// 				x := i*tileSize + (i+1)*tileMargin
// 				y := j*tileSize + (j+1)*tileMargin
// 				op.GeoM.Translate(float64(x), float64(y))
// 				r, g, b, a := colorToScale(tileBackgroundColor(v))
// 				op.ColorM.Scale(r, g, b, a)
// 				tileImage, _ := ebiten.NewImage(tileSize, tileSize, ebiten.FilterNearest)
// 				tileImage.Fill(color.White)
// 				image.DrawImage(tileImage, op)
// 			}
//
// 		}
// 	}
// }
