package main

import (
	_ "embed"
	"fmt"
	"github.com/Stexxe/bc/internal/app/anim"
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

var tankAnim = "up"
var tankAnims = map[string]anim.Descriptor{
	"up":    {X: 1, Y: 130, FramesCount: 2},
	"left":  {X: 34, Y: 129, FramesCount: 2},
	"down":  {X: 65, Y: 129, FramesCount: 2},
	"right": {X: 97, Y: 129, FramesCount: 2},
}

var pressedKey sdl.Keycode = sdl.K_UNKNOWN

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("BATTLE CITY", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
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

	size := int32(13)
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
				if e.Type == sdl.KEYDOWN && pressedKey == sdl.K_UNKNOWN {
					speed = tankSpeed
					pressedKey = e.Keysym.Sym

					switch pressedKey {
					case sdl.K_w:
						direction = util.VectorUp
						tankAnim = "up"
					case sdl.K_a:
						direction = util.VectorLeft
						tankAnim = "left"
					case sdl.K_s:
						direction = util.VectorDown
						tankAnim = "down"
					case sdl.K_d:
						direction = util.VectorRight
						tankAnim = "right"
					}
				} else if e.Type == sdl.KEYUP && pressedKey == e.Keysym.Sym {
					fmt.Println(sdl.GetKeyName(sdl.Keycode(pressedKey)), sdl.GetKeyName(sdl.Keycode(e.Keysym.Sym)))
					pressedKey = sdl.K_UNKNOWN
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

		ad, _ := tankAnims[tankAnim]
		if curFrame >= ad.FramesCount {
			curFrame = 0
		}

		tankPos = tankPos.Sum(direction.MulScalar(int32(speed * float32(delta))))

		// Draw
		renderer.SetDrawColor(0, 0, 1, 255)
		renderer.Clear()

		x := ad.X
		y := ad.Y
		if speed > 0 {
			x += curFrame * (size + gap)
		}

		renderer.Copy(texture, &sdl.Rect{x, y, size, size}, &sdl.Rect{tankPos.X, tankPos.Y, size * scale, size * scale})
		renderer.Present()

		curTime = sdl.GetTicks64()
		sdl.Delay(16)
	}
}
