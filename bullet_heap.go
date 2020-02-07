package main

import "github.com/faiface/pixel/imdraw"

// BulletHeap collection of active bullets that should be tracked and added to the screen
// this isn't a 'real' heap in the classical programming sense.  i'm just bad at naming things
type BulletHeap struct {
	Bullets []Bullet
}

// Count Returns number of active bullets in heap
func (bh *BulletHeap) Count() int {
	return len(bh.Bullets)
}

// Add Add bullets to heap
func (bh *BulletHeap) Add(b Bullet) {
	bh.Bullets = append(bh.Bullets, b)
}

// Draw applies drawings to given imgdraw.IMDraw and cleans up expired bullets
func (bh *BulletHeap) Draw(imd *imdraw.IMDraw) {
	cleanBullets := []int{}

	for i, b := range bh.Bullets {
		if b.IsDone() {
			cleanBullets = append(cleanBullets, i)
		} else {
			b.Update()
			b.Draw(imd)
		}
	}
	if len(cleanBullets) > 0 {
		for _, i := range cleanBullets {
			bh.Bullets = append(bh.Bullets[:i], bh.Bullets[i+1:]...)
		}
	}

}
