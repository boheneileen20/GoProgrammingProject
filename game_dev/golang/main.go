package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenHeight = 600
	screenWidth= 600
)

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil{
		fmt.Println("Initializing SDL:",err)
		return
	}

	//CREATES THE WINDOW FOR THE GAME
	window, err := sdl.CreateWindow(
		"SDL all set",
		sdl.WINDOWPOS_UNDEFINED,sdl.WINDOWPOS_UNDEFINED,screenWidth,screenHeight,
		sdl.WINDOW_OPENGL)

		if err != nil {
			fmt.Println("Initializing window:",err)
			return
		}

		defer window.Destroy()

		//HELPS US DRAWING THINGS ON THE WINDOW
	renderer,err := sdl.CreateRenderer(window,-1,sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Println("Initializing renderer:",err)
		return
	}

	defer renderer.Destroy()


	for{
		//CREATES A COLOR OF WHITE
		renderer.SetDrawColor(255,255,255,255)
		renderer.Clear()
		//HELPS US TO SHOW THE COLOR
		renderer,Present()
	}


}
