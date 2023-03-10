package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	// hello-worldを表示させるファイルをインポート
	"github.com/O6lvl4/sandbox-ebitengine/hello"
)

func main() {
	// helloパッケージで宣言されたGameをgameへ格納
	game, err := hello.NewGame()
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	// helloパッケージで宣言されているGameを実行する
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
