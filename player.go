package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const playerSize = 16.0

type Player struct {
	Pos      pixel.Vec
	Color    color.RGBA
	showline bool
}

func (p *Player) Draw(imd *imdraw.IMDraw) {
	imd.Color = p.Color
	imd.Push(pixel.V(p.Pos.X, p.Pos.Y), pixel.V(p.Pos.X+playerSize, p.Pos.Y+playerSize))
	imd.Rectangle(0)
	imd.Color = colornames.Black
	imd.Push(pixel.V(p.Pos.X, p.Pos.Y), pixel.V(p.Pos.X+playerSize, p.Pos.Y+playerSize))
	imd.Rectangle(1)
}

func (p *Player) Center() pixel.Vec {
	return pixel.V(p.Pos.X+(playerSize/2), p.Pos.Y+(playerSize/2))
}

func (p *Player) Translate(v pixel.Vec) {
	p.Pos.X += v.X
	p.Pos.Y += v.Y
}
