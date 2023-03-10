package hello

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Hello, WorldとTPSを画面に表示させる
	ebitenutil.DebugPrint(screen, "Hello, World!\n")
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nTPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// main.goやmobile.goで、ここで作ったGame構造体を使えるようにする重要な部分
func NewGame() (*Game, error) {
	game := &Game{}
	return game, nil
}
