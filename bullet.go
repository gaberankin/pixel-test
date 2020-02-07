package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Bullet interface {
	Update()
	Draw(imd *imdraw.IMDraw)
	IsDone() bool
}

type StraightBullet struct {
	win     *pixelgl.Window
	Pos     pixel.Vec
	Vector  pixel.Vec
	Speed   float64
	TimeEnd time.Time
}

func NewStraightBullet(win *pixelgl.Window, pos, vector pixel.Vec, lifetime time.Duration) *StraightBullet {
	return &StraightBullet{
		win:     win,
		Pos:     pos,
		Vector:  vector,
		TimeEnd: time.Now().Add(lifetime),
		Speed:   5.0,
	}
}

func (b *StraightBullet) Update() {
	x, y := b.Pos.XY()
	vx, vy := b.Vector.XY()
	b.Pos.X = x + (vx * b.Speed)
	b.Pos.Y = y + (vy * b.Speed)
}

func (b *StraightBullet) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Black
	imd.Push(b.Pos)
	imd.Circle(2.0, 0)
	imd.Color = colornames.White
	imd.Push(b.Pos)
	imd.Circle(1.0, 0)
}

func (b *StraightBullet) IsDone() bool {
	return !b.win.Bounds().Contains(b.Pos) || time.Now().After(b.TimeEnd)
}
