package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/kage-desk/display"
)

//go:embed image.kage
var imageShaderProgram []byte

//go:embed circle.kage
var shaderProgram []byte

func main() {
	fmt.Println("Hello world!")
	// shader, err := ebiten.NewShader(shaderProgram)
	//
	shader, err := ebiten.NewShader(imageShaderProgram)
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{shader: shader, angle: 80.0}
	ebiten.SetWindowTitle("Shaders test")
	bounds := display.ImageSpiderCatDog().Bounds()
	ebiten.SetWindowSize(bounds.Dx(), bounds.Dy()*2)

	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	shader   *ebiten.Shader
	vertices [4]ebiten.Vertex
	angle    int
}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{33, 66, 99, 255})
	g.drawImageWithReflection(screen)
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	bounds := display.ImageSpiderCatDog().Bounds()
	return bounds.Dx(), bounds.Dy() * 2
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	g.angle = g.angle + 1
	if g.angle >= 360 {
		g.angle = 0
	}

	return nil
}
