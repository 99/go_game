package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)
const (
	screenWidth = 600
	screenHeight = 800
)
func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}
	window, err := sdl.CreateWindow(
		"My first Go Game",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Init window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Init renderer:", err)
		return
	}
	defer renderer.Destroy()

	plr, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("creating a player:", err)
		return
	}

	var enemies []enemy
	for i:=0; i <5; i++ {
		for j:=0; j<3; j++ {
			x := (float64(i) / 5) * screenWidth + (enemySize/2.0)
			y := float64(j) * enemySize +(enemySize/2.0)

			enemy, err := newEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("creating a enemy:", err)
				return
			}
			enemies = append(enemies, enemy)
		}
	}


		renderer.Present()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			//Command +W
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		plr.draw(renderer)
		plr.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}
		renderer.Present()
	}
}