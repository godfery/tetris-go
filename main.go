package main

import (
	"game/tetris"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/pkg/profile"
)

var gggg *tetris.GameUnit

func draw() {
	gggg.Draw()

	// animatingTiles := map[*Tile]struct{}{}
	// nonAnimatingTiles := map[*Tile]struct{}{}
	// for t := range b.tiles {
	// 	if t.IsMoving() {
	// 		animatingTiles[t] = struct{}{}
	// 	} else {
	// 		nonAnimatingTiles[t] = struct{}{}
	// 	}
	// }
	// for t := range nonAnimatingTiles {
	// 	t.Draw(boardImage)
	// }
	// for t := range animatingTiles {
	// 	t.Draw(boardImage)
	// }
}

func update(screen *ebiten.Image) error {

	if ebiten.IsRunningSlowly() {
		return nil
	}

	gggg.Update(screen)

	// fmt.Println(ebiten.IsKeyPressed(ebiten.KeyRight))
	// game.Draw(screen)
	return nil
}

func main() {
	// aaa.Out()
	// defer profile.Start().Stop()
	defer profile.Start(profile.MemProfileRate(2048)).Stop()
	var err error
	// game, err = aaa.NewGame()
	gggg, err = tetris.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	// ebiten.SetFullscreen(true)
	if err := ebiten.Run(update, tetris.ScreenWidth, tetris.ScreenHeight, 2, "2048 (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
