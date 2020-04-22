package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const enemySize = 105

type enemy struct {
	tex * sdl.Texture
	x,y float64
}

func newEnemy(renderer *sdl.Renderer, x, y float64)(en enemy, err error) {
	img, err := sdl.LoadBMP("sprite/enemy.bmp")
	if err != nil {
		return enemy{}, fmt.Errorf("loading enemy sprite: %v" , err)
	}
	defer img.Free()
	en.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return enemy{}, fmt.Errorf("Creating enemy texture: %v" , err)
	}
	en.x = x
	en.y = y
	return en, nil
}

func (en *enemy) draw(renderer *sdl.Renderer) {
	x := en.x - enemySize/2.0
	y := en.y - enemySize/2.0
	renderer.CopyEx(en.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H:105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H:105},
		180,
		&sdl.Point{X: enemySize/2, Y: enemySize/2},
		sdl.FLIP_NONE)
}