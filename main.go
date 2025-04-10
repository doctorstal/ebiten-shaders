package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"log"
	"math"

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

func (g *Game) drawImageWithReflection(screen *ebiten.Image) {
	// map the vertices to the target image
	bounds := screen.Bounds()
	g.vertices[0].DstX = float32(bounds.Min.X) // top-left
	g.vertices[0].DstY = float32(bounds.Min.Y) // top-left
	g.vertices[1].DstX = float32(bounds.Max.X) // top-right
	g.vertices[1].DstY = float32(bounds.Min.Y) // top-right
	g.vertices[2].DstX = float32(bounds.Min.X) // bottom-left
	g.vertices[2].DstY = float32(bounds.Max.Y) // bottom-left
	g.vertices[3].DstX = float32(bounds.Max.X) // bottom-right
	g.vertices[3].DstY = float32(bounds.Max.Y) // bottom-right

	// set the source image sampling coordinates
	srcBounds := display.ImageSpiderCatDog().Bounds()
	g.vertices[0].SrcX = float32(srcBounds.Min.X) // top-left
	g.vertices[0].SrcY = float32(srcBounds.Min.Y) // top-left
	g.vertices[1].SrcX = float32(srcBounds.Max.X) // top-right
	g.vertices[1].SrcY = float32(srcBounds.Min.Y) // top-right
	g.vertices[2].SrcX = float32(srcBounds.Min.X) // bottom-left
	g.vertices[2].SrcY = float32(srcBounds.Max.Y) // bottom-left
	g.vertices[3].SrcX = float32(srcBounds.Max.X) // bottom-right
	g.vertices[3].SrcY = float32(srcBounds.Max.Y) // bottom-right

	// triangle shader options
	var shaderOpts ebiten.DrawTrianglesShaderOptions
	shaderOpts.Images[0] = display.ImageSpiderCatDog()
	shaderOpts.Uniforms = make(map[string]any)
	shaderOpts.Uniforms["MirrorAlphaMult"] = float32(0.2)
	shaderOpts.Uniforms["VertDisplacement"] = 28

	// draw shader
	indices := []uint16{0, 1, 2, 2, 1, 3} // map vertices to triangles
	screen.DrawTrianglesShader(g.vertices[:], indices, g.shader, &shaderOpts)
}

func (g *Game) drawImage(screen *ebiten.Image) {
	bounds := screen.Bounds()
	g.vertices[0].DstX = float32(bounds.Min.X) // top-left
	g.vertices[0].DstY = float32(bounds.Min.Y)
	g.vertices[1].DstX = float32(bounds.Max.X) // top-right
	g.vertices[1].DstY = float32(bounds.Min.Y)
	g.vertices[2].DstX = float32(bounds.Min.X) // bottom-left
	g.vertices[2].DstY = float32(bounds.Max.Y)
	g.vertices[3].DstX = float32(bounds.Max.X) // botton-right
	g.vertices[3].DstY = float32(bounds.Max.Y)

	// set the source image sampling coordinates
	srcBounds := display.ImageSpiderCatDog().Bounds()
	g.vertices[0].SrcX = float32(srcBounds.Min.X) // top-left
	g.vertices[0].SrcY = float32(srcBounds.Min.Y) // top-left
	g.vertices[1].SrcX = float32(srcBounds.Max.X) // top-right
	g.vertices[1].SrcY = float32(srcBounds.Min.Y) // top-right
	g.vertices[2].SrcX = float32(srcBounds.Min.X) // bottom-left
	g.vertices[2].SrcY = float32(srcBounds.Max.Y) // bottom-left
	g.vertices[3].SrcX = float32(srcBounds.Max.X) // bottom-right
	g.vertices[3].SrcY = float32(srcBounds.Max.Y) // bottom-right

	var shaderOpts ebiten.DrawTrianglesShaderOptions

	shaderOpts.Uniforms = make(map[string]any)
	shaderOpts.Images[0] = display.ImageSpiderCatDog()

	indicies := []uint16{0, 1, 2, 2, 1, 3}
	screen.DrawTrianglesShader(g.vertices[:], indicies, g.shader, &shaderOpts)
}

func (g *Game) drawCircle(screen *ebiten.Image) {
	bounds := screen.Bounds()
	g.vertices[0].DstX = float32(bounds.Min.X) // top-left
	g.vertices[0].DstY = float32(bounds.Min.Y)
	g.vertices[1].DstX = float32(bounds.Max.X) // top-right
	g.vertices[1].DstY = float32(bounds.Min.Y)
	g.vertices[2].DstX = float32(bounds.Min.X) // bottom-left
	g.vertices[2].DstY = float32(bounds.Max.Y)
	g.vertices[3].DstX = float32(bounds.Max.X) // botton-right
	g.vertices[3].DstY = float32(bounds.Max.Y)

	var shaderOpts ebiten.DrawTrianglesShaderOptions

	shaderOpts.Uniforms = make(map[string]any)
	shaderOpts.Uniforms["Center"] = []float32{
		float32(screen.Bounds().Dx()) / 2,
		float32(screen.Bounds().Dy()) / 2,
	}
	shaderOpts.Uniforms["Radius"] = float32(60.0 + 20.0*math.Sin(float64(g.angle)*math.Pi/180.0))

	indicies := []uint16{0, 1, 2, 2, 1, 3}
	screen.DrawTrianglesShader(g.vertices[:], indicies, g.shader, &shaderOpts)
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
