package game

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Width  int32
	Height int32
	Title  string

	TargetFps       int32
	BackgroundColor raylib.Color
}

func CreateGame(width, height int32, title string, bgColor raylib.Color) *Game {
	return &Game{
		Width:           width,
		Height:          height,
		Title:           title,
		TargetFps:       60,
		BackgroundColor: bgColor,
	}
}

func (g *Game) Run() {
	raylib.InitWindow(g.Width, g.Height, g.Title)
	raylib.SetTargetFPS(g.TargetFps)

	leftPaddle := CreatePaddle(755, float32(g.Height/3), 20, 220, raylib.KeyUp, raylib.KeyDown)
	rightPaddle := CreatePaddle(20, float32(g.Height/3), 20, 220, raylib.KeyW, raylib.KeyS)

	ball := CreateBall(300, 300, 10, g.Width/2, g.Height/2, leftPaddle, rightPaddle)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()

		raylib.ClearBackground(g.BackgroundColor)

		if raylib.IsKeyDown(raylib.KeyR) {
			leftPaddle.SetPaddlePos(755, float32(g.Height/3))
			rightPaddle.SetPaddlePos(20, float32(g.Height/3))
			ball.SetBallPos(g.Width/2, g.Height/2, 300, 300)
		}

		leftPaddle.Update()
		rightPaddle.Update()
		ball.Update()

		raylib.EndDrawing()
	}

	raylib.CloseWindow()
}
