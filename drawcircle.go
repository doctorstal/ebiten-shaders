package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

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
