package main

import (
	_ "embed"
	"github.com/Stexxe/bc/internal/app/util"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

//go:embed assets/spritesheet.png
var spriteSheet []byte

var speed = float32(0)
var direction = util.VectorUp
var tankPos = util.NewVector(100, 100)

const tankSpeed = 0.1

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
			switch e := event.(type) {
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					speed = tankSpeed

					switch e.Keysym.Sym {
					case sdl.K_w:
						direction = util.VectorUp
					case sdl.K_a:
						direction = util.VectorLeft
					case sdl.K_s:
						direction = util.VectorDown
					case sdl.K_d:
						direction = util.VectorRight
					}
				} else if e.Type == sdl.KEYUP {
					speed = 0
				}
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

		tankPos = tankPos.Sum(direction.MulScalar(int32(speed * float32(delta))))

		// Draw
		renderer.SetDrawColor(169, 169, 169, 255)
		renderer.Clear()

		x := ssTankX + curFrame*(size+gap)
		y := ssTankY

		renderer.Copy(texture, &sdl.Rect{x, y, size, size}, &sdl.Rect{tankPos.X, tankPos.Y, size * scale, size * scale})
		renderer.Present()

		curTime = sdl.GetTicks64()
		sdl.Delay(16)
	}
}
