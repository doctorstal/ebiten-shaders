package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/kage-desk/display"
)

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
