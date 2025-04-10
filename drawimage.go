package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/kage-desk/display"
)

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
