package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const(
	screenWidth=600
	screenHeight=800
)

func main(){
	if err:=sdl.Init(sdl.INIT_EVERYTHING);err !=nil{
		fmt.Println("intializing SDL:",err)
		return
	}
	
	window,err:=sdl.CreateWindow(
		"Gaming In Go",
		sdl.WINDOWPOS_UNDEFINED,sdl.WINDOWPOS_UNDEFINED,
		screenWidth,screenHeight,
		sdl.WINDOW_OPENGL)
	if err!=nil{
		fmt.Println("intializing renderer",err)
		return
	}
	defer window.Destroy()
	
	renderer,err:=sdl.CreateRenderer(window,-1,sdl.RENDERER_ACCELERATED)
	if err!=nil{
		fmt.Println("intializing renderer:",err)
		return
	}
	defer renderer.Destroy()
	
	plr,err:=newPlayer(renderer)
	if err!=nil{
		fmt.Println("creating new player:",err)
	}
	
	img,err:=sdl.LoadBMP("Sprites/player.bmp")
	if err!=nil{
		fmt.Println("loading player sprite:",err)
		return
	}
	defer img.Free()
	
	playerTex,err:=renderer.CreateTextureFromSurface(img)
	if err!=nil{
		fmt.Println("creating player texture:",err)
		return
	}
	
	defer playerTex.Destroy()
	
	for{
		for event:=sdl.PollEvent();event!=nil;event = sdl.PollEvent(){
			switch event.(type){
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255,255,255,255)
		renderer.Clear()
		
		plr.draw(renderer)
		
		renderer.Present()
	}
}