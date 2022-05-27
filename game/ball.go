package game

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	xVelocity float32
	yVelocity float32

	radius float32

	xPos int32
	yPos int32

	leftPaddle  *Paddle
	rightPaddle *Paddle
}

func CreateBall(xVelocity, yVelocity, radius float32, xPos, yPos int32, left, right *Paddle) *Ball {
	return &Ball{
		xVelocity:   xVelocity,
		yVelocity:   yVelocity,
		xPos:        xPos,
		yPos:        yPos,
		radius:      radius,
		leftPaddle:  left,
		rightPaddle: right,
	}
}

func (b *Ball) SetBallPos(xPos, yPos int32, xVelocity, yVelocity float32) {
	b.xPos = xPos
	b.yPos = yPos

	b.xVelocity = xVelocity
	b.yVelocity = yVelocity
}

func (b *Ball) Update() {
	b.xPos += int32(b.xVelocity * raylib.GetFrameTime())
	b.yPos += int32(b.yVelocity * raylib.GetFrameTime())

	if b.yPos < 0 {
		b.yPos = 0
		b.yVelocity = -b.yVelocity
	} else if b.yPos > int32(raylib.GetScreenHeight()) {
		b.yVelocity = -b.yVelocity
	}

	if raylib.CheckCollisionCircleRec(
		raylib.Vector2{X: float32(b.xPos), Y: float32(b.yPos)}, b.radius, b.leftPaddle.GetPaddleRect(),
	) {
		b.xVelocity *= -1.1
	} else if raylib.CheckCollisionCircleRec(
		raylib.Vector2{X: float32(b.xPos), Y: float32(b.yPos)}, b.radius, b.rightPaddle.GetPaddleRect(),
	) {
		b.xVelocity *= -1.1
	}

	raylib.DrawCircle(b.xPos, b.yPos, b.radius, raylib.White)
}
