package main

import (
	"pong/game"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	title        = "Pong"
)

func init() {
	raylib.SetConfigFlags(raylib.FlagVsyncHint)
	raylib.SetConfigFlags(raylib.FlagMsaa4xHint)
}

func main() {
	g := game.CreateGame(screenWidth, screenHeight, title, raylib.Black)

	g.Run()
}
