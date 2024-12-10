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

	//surface, err := window.GetSurface()
	//if err != nil {
	//	panic(err)
	//}

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
		log.Fatalf("Failed to create texture: %s\n", err)
	}
	defer texture.Destroy()

	// Clear the renderer and draw the texture
	renderer.Clear()
	surface.FillRect(nil, sdl.Color{240, 240, 240, 255}.Uint32())
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	//rect := sdl.Rect{0, 0, 200, 200}
	//colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	//pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	//surface.FillRect(&rect, pixel)
	//window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		sdl.Delay(33)
	}
}
