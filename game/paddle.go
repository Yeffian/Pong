package game

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	xPos float32
	yPos float32

	upKey   int32
	downKey int32

	height float32
	width  float32
}

func CreatePaddle(xPos, yPos, width, height float32, upKey, downKey int32) *Paddle {
	return &Paddle{
		xPos:    xPos,
		yPos:    yPos,
		upKey:   upKey,
		downKey: downKey,
		width:   width,
		height:  height,
	}
}

func (p *Paddle) GetPaddleRect() raylib.Rectangle {
	return raylib.Rectangle{
		X:      p.xPos,
		Y:      p.yPos,
		Width:  p.width,
		Height: p.height,
	}
}

func (p *Paddle) Update() {
	if raylib.IsKeyDown(p.upKey) {
		p.yPos -= 5
	} else if raylib.IsKeyDown(p.downKey) {
		p.yPos += 5
	}

	raylib.DrawRectangleRec(p.GetPaddleRect(), raylib.White)
}

func (p *Paddle) SetPaddlePos(xPos, yPos float32) {
	p.xPos = xPos
	p.yPos = yPos
}
