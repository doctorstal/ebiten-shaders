package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed circle.kage
var shaderProgram []byte

func main() {
	fmt.Println("Hello world!")
	shader, err := ebiten.NewShader(shaderProgram)

	game := &Game{shader: shader, angle: 80.0}
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowTitle("Shaders test")
	ebiten.SetWindowSize(512, 512)

	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	shader    *ebiten.Shader
	verticies [4]ebiten.Vertex
	angle     int
}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	bounds := screen.Bounds()
	g.verticies[0].DstX = float32(bounds.Min.X) // top-left
	g.verticies[0].DstY = float32(bounds.Min.Y)
	g.verticies[1].DstX = float32(bounds.Max.X) // top-right
	g.verticies[1].DstY = float32(bounds.Min.Y)
	g.verticies[2].DstX = float32(bounds.Min.X) // bottom-left
	g.verticies[2].DstY = float32(bounds.Max.Y)
	g.verticies[3].DstX = float32(bounds.Max.X) // botton-right
	g.verticies[3].DstY = float32(bounds.Max.Y)

	var shaderOpts ebiten.DrawTrianglesShaderOptions

	shaderOpts.Uniforms = make(map[string]any)
	shaderOpts.Uniforms["Center"] = []float32{
		float32(screen.Bounds().Dx()) / 2,
		float32(screen.Bounds().Dy()) / 2,
	}
	shaderOpts.Uniforms["Radius"] = float32(60.0 + 20.0*math.Sin(float64(g.angle)*math.Pi/180.0))

	indicies := []uint16{0, 1, 2, 2, 1, 3}
	screen.DrawTrianglesShader(g.verticies[:], indicies, g.shader, &shaderOpts)
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 512, 512
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	g.angle = g.angle + 1
	if g.angle >= 360 {
		g.angle = 0
	}
	return nil
}
