package main 

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)


type player struct{

	tex *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (p player, err error){

	
	//image loaded
	img, err := sdl.LoadBMP("./f.bmp")
    if err != nil {
		return player{}, fmt.Errorf("loading player gopher: %v",err)
		
	}
    defer img.Free()
	// A texture to process the image that is loaded on the screen

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("loading player texture: %v",err)
		
	}
	
  return p, nil
}


func (p *player) draw(renderer *sdl.Renderer){

	//renderer copy helps to load the img which converted to texture
	    renderer.Copy(p.tex, 
		&sdl.Rect{X:0,Y:0, W:105, H:105},
		&sdl.Rect{X:0,Y:0, W:70, H:50})
}