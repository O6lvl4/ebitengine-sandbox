package mobile

import (
	"github.com/O6lvl4/sandbox-ebitengine/hello"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	// helloパッケージで宣言されたGameをgameへ格納
	game, err := hello.NewGame()
	if err != nil {
		panic(err)
	}

	// helloパッケージで宣言されているGameをmobile用にセットする
	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
