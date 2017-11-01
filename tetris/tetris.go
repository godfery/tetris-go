package tetris

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// object 方块定义
type area struct {
	param [4][4]int
	xLen  int
	yLen  int
}

//方块管理对象
type TetrisObject struct {
	area        area
	posionX     int
	posionY     int
	currentTime int64
	tetrisPos   int
	LineCage    [10][10]int
	currentType int
	index       int
}

// 方块集合
type areas struct {
	all      [4]area
	rotation [4][4]area
}

//定义变量
var allAreas areas
var loopEnd bool

func init() {
	loopEnd = false
	allAreas = InitCage()
}

func NewTetris() *TetrisObject {
	var a [10][10]int
	tetrisObject := &TetrisObject{allAreas.all[0], 3, 0, 0, size, a, 0, 0}
	return tetrisObject

}
func (tetrisObject *TetrisObject) randTetrisShow() {
	tetrisObject.area = allAreas.all[tetrisObject.currentType]
	tetrisObject.posionX = 3
	tetrisObject.posionY = 0
	loopEnd = false

}

//移动到下一格 碰撞检验

// 如果有数据的话 直接落到上面
func (tetrisObject *TetrisObject) fadCage() {
	fmt.Println("fadcage")
	var tt [size][size]int
	start := size - 1
	for xx := size - 1; xx >= tetrisObject.tetrisPos; xx-- {
		sum := 0
		for jj := 0; jj < size; jj++ {
			sum = sum + tetrisObject.LineCage[xx][jj]
		}
		fmt.Println(sum, size)
		if sum != size {

			tt[start] = tetrisObject.LineCage[xx]
			start--
		}
	}
	tetrisObject.LineCage = tt
}

// 方块下落完毕，放到lineCage集合中，以便保存
func (tetrisObject *TetrisObject) addTetrisToLindCage() {
	// fmt.Println("addTetrisToLindCage")
	defer tetrisObject.fadCage()
	area := tetrisObject.area.param

	for xx := 0; xx < 4; xx++ {
		for jj := 0; jj < 4; jj++ {

			if area[xx][jj] > 0 {
				j := tetrisObject.posionY + xx
				i := tetrisObject.posionX + jj
				if j < tetrisObject.tetrisPos {
					tetrisObject.tetrisPos = j
				}
				// fmt.Println(i, j, xx, jj)
				tetrisObject.LineCage[j][i] = area[xx][jj]
			}

		}
	}

}

//机器人 随机移动 测试用的
func (to *TetrisObject) robotRandMove() {
	if time.Now().UnixNano()/1e6-to.currentTime > 800 && !loopEnd {
		to.currentTime = time.Now().UnixNano() / 1e6
		to.posionY = to.posionY + 1

	}
}

//根据 按键操作
func (to *TetrisObject) MoveByDir(i *Input) {

	d, _ := i.Dir()
	switch d {
	case 2:

		// rand := Generate_Randnum(3)
		// fmt.Println(rand)
		to.index++
		if to.index > 3 {
			to.index = 0
		}
		to.area = allAreas.rotation[to.currentType][to.index]
	case 3:

		to.posionX = to.posionX - 1
		if to.posionX < 0 {
			to.posionX = 0
		}
		break
	case 1:

		if to.posionX+to.area.xLen < size {
			to.posionX = to.posionX + 1
		}

		break
	}
}

// 画图像
func (to *TetrisObject) Draw(image *ebiten.Image) {

	to.robotRandMove()
	// fmt.Println(time.Now().UnixNano() - to.currentTime)
	for xx := 0; xx < 4; xx++ {
		for jj := 0; jj < 4; jj++ {
			area := to.area.param
			if area[xx][jj] > 0 {
				j := to.posionY + xx
				i := to.posionX + jj
				if i >= size {
					i = size - 1
				}
				if i < 0 {
					i = 0
				}
				v := 4096

				if j >= size-1 || to.LineCage[j+1][i] == 1 {
					loopEnd = true
					to.addTetrisToLindCage()
					to.randTetrisShow()
				}
				op := &ebiten.DrawImageOptions{}
				x := i*tileSize + (i+1)*tileMargin
				y := j*tileSize + (j+1)*tileMargin
				op.GeoM.Translate(float64(x), float64(y))
				r, g, b, a := colorToScale(tileBackgroundColor(v))
				op.ColorM.Scale(r, g, b, a)
				tileImage, _ := ebiten.NewImage(tileSize, tileSize, ebiten.FilterNearest)
				tileImage.Fill(color.White)
				image.DrawImage(tileImage, op)
			}

		}
	}

}
