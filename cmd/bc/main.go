package main

import (
	_ "embed"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

//go:embed assets/spritesheet.png
var spriteSheet []byte

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("PACMAN", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	if err := img.Init(img.INIT_PNG); err != nil {
		log.Fatal(err)
	}

	defer img.Quit()

	rwops, err := sdl.RWFromMem(spriteSheet)

	if err != nil {
		log.Fatal(err)
	}

	surface, err := img.LoadRW(rwops, true)
	if err != nil {
		log.Fatal(err)
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal(err)
	}
	defer texture.Destroy()

	ssTankX := int32(1)
	ssTankY := int32(130)
	size := int32(13)
	frames := int32(2)
	curFrame := int32(0)
	scale := int32(2)
	gap := int32(3)
	frameDur := uint64(150)
	accTime := uint64(0)

	running := true
	var delta uint64
	curTime := sdl.GetTicks64()

	for running {
		delta = sdl.GetTicks64() - curTime

		// Events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		// State
		accTime += delta
		if accTime >= frameDur {
			accTime = 0
			curFrame++
		}

		if curFrame >= frames {
			curFrame = 0
		}

		// Draw
		renderer.SetDrawColor(169, 169, 169, 255)
		renderer.Clear()

		x := ssTankX + curFrame*(size+gap)
		y := ssTankY

		renderer.Copy(texture, &sdl.Rect{x, y, size, size}, &sdl.Rect{100, 100, size * scale, size * scale})
		renderer.Present()

		curTime = sdl.GetTicks64()
		sdl.Delay(16)
	}
}
