package main

import (
	"fmt"

	_ "embed"

	"github.com/tinne26/kage-desk/display"
)

//go:embed shader.kage
var shaderProgram []byte

func main() {
	fmt.Println("Hello world!")
	display.SetTitle("gradient")
	display.SetSize(512, 512)
	display.Shader(shaderProgram)
}
