package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const playerSpeed =0.5

type player struct {
	tex *sdl.Texture
	x,y float64
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	img, err := sdl.LoadBMP("Sprites/player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer img.Free()
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("creating player texture: %v", err)
	}

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 105, H: 105})
}
func (p *player) update(){
	keys:=sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT]==1{
		p.x -= playerSpeed
	}else if keys[sdl.SCANCODE_RIGHT]==1{
		p.x += playerSpeed
	}
}