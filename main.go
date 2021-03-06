package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// screen size.  shooting to make this configurable
const WIDTH = 1024
const HEIGHT = 768

// draws a white line with a black outline.  using this for debugging right now.
func drawLine(imd *imdraw.IMDraw, start pixel.Vec, end pixel.Vec, alpha float64) {
	imd.SetColorMask(pixel.Alpha(alpha))
	imd.Color = colornames.Black
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(start, end)
	imd.Line(3)
	imd.EndShape = imdraw.NoEndShape
	imd.Color = colornames.White
	imd.Push(start, end)
	imd.Line(1)
	imd.SetColorMask(pixel.Alpha(1))
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Test 2",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	p := Player{Pos: pixel.V(cfg.Bounds.W()/2, cfg.Bounds.H()/2), Color: colornames.Red}

	currentAlpha := 1.0

	bHeap := &BulletHeap{}
	debugTicker := time.Tick(time.Second)
	frames := 0
	playerSpeed := 3.0

	for !win.Closed() {
		pCenter := p.Center()
		mPos := win.MousePosition()
		winWidth, winHeight := win.Bounds().Size().XY()

		deltaX := 0.0
		deltaY := 0.0

		// controls are WASD for directions and mouse clicks to fire in direction of mouse pointer.
		if win.Pressed(pixelgl.KeyS) && (pCenter.Y-playerSpeed) > 0 {
			deltaY = playerSpeed * -1.0
		}
		if win.Pressed(pixelgl.KeyW) && (pCenter.Y+playerSpeed) < winHeight {
			deltaY = playerSpeed
		}
		if win.Pressed(pixelgl.KeyA) && (pCenter.X-playerSpeed) > 0 {
			deltaX = playerSpeed * -1.0
		}
		if win.Pressed(pixelgl.KeyD) && (pCenter.X+playerSpeed) < winWidth {
			deltaX = playerSpeed
		}
		p.Translate(pixel.V(deltaX, deltaY))

		if win.Pressed(pixelgl.MouseButtonLeft) {
			currentAlpha = 0.2
		} else {
			currentAlpha = 1.0
		}
		// TODO - make this fire bullets on a timer
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			bHeap.Add(NewStraightBullet(
				win,
				pCenter,
				pixel.V(mPos.X-pCenter.X, mPos.Y-pCenter.Y).Unit(),
				3*time.Second,
			))
		} else if win.JustReleased(pixelgl.MouseButtonLeft) {

		}

		imd.Clear()

		drawLine(imd, pCenter, mPos, currentAlpha)

		bHeap.Draw(imd)

		p.Draw(imd)

		win.Clear(colornames.Gray)
		imd.Draw(win)
		win.Update()

		frames++
		select {
		case <-debugTicker:
			win.SetTitle(fmt.Sprintf("%s | fps: %d | Angle: %.02f | Bullets: %d | Player: (%.02f, %.02f)", cfg.Title, frames, pointAngle(mPos, pCenter), bHeap.Count(), p.Center().X, p.Center().Y))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
