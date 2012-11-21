// W3C is a set of selected colors by the W3Consortium.

package main

import (
	"image/color"
)

const (
	aqua = iota
	aquamarine
	blueviolet
	brown
	burlywood
	cadetblue
	chartreuse
	chocolate
	coral
	cornflowerblue
	crimson
	darkblue
	darkcyan
	darkgoldenrod
	darkgray
	darkgreen
	darkkhaki
	darkmagenta
	darkolivegreen
	darkorange
	darkorchid
	darkred
	darksalmon
	darkseagreen
	darkslateblue
	darkslategray
	darkslategrey
	darkturquoise
	darkviolet
	deeppink
	deepskyblue
	dimgray
	dodgerblue
	firebrick
	forestgreen
	fuchsia
	gold
	goldenrod
	green
	greenyellow
	hotpink
	indianred
	indigo
	khaki
	lawngreen
	lightblue
	lightcoral
	lightgreen
	lightpink
	lightsalmon
	lightseagreen
	lightskyblue
	lightslategray
	lightsteelblue
	lime
	limegreen
	magenta
	maroon
	mediumaquamarine
	mediumblue
	mediumorchid
	mediumpurple
	mediumseagreen
	mediumslateblue
	mediumspringgreen
	mediumturquoise
	mediumvioletred
	midnightblue
	olive
	olivedrab
	orange
	orangered
	orchid
	palegoldenrod
	palegreen
	paleturquoise
	palevioletred
	peachpuff
	peru
	pink
	plum
	powderblue
	purple
	red
	rosybrown
	royalblue
	saddlebrown
	salmon
	sandybrown
	seagreen
	sienna
	silver
	skyblue
	slateblue
	slategray
	springgreen
	steelblue
	tan
	teal
	thistle
	tomato
	turquoise
	violet
	wheat
	yellow
	yellowgreen
)

var W3C = color.Palette{
	color.RGBA{ 0, 255, 255, 255},
	color.RGBA{127, 255, 212, 255},
	color.RGBA{138, 43, 226, 255},
	color.RGBA{165, 42, 42, 255},
	color.RGBA{222, 184, 135, 255},
	color.RGBA{ 95, 158, 160, 255},
	color.RGBA{127, 255, 0, 255},
	color.RGBA{210, 105, 30, 255},
	color.RGBA{255, 127, 80, 255},
	color.RGBA{100, 149, 237, 255},
	color.RGBA{220, 20, 60, 255},
	color.RGBA{ 0, 0, 139, 255},
	color.RGBA{ 0, 139, 139, 255},
	color.RGBA{184, 134, 11, 255},
	color.RGBA{169, 169, 169, 255},
	color.RGBA{ 0, 100, 0, 255},
	color.RGBA{189, 183, 107, 255},
	color.RGBA{139, 0, 139, 255},
	color.RGBA{ 85, 107, 47, 255},
	color.RGBA{255, 140, 0, 255},
	color.RGBA{153, 50, 204, 255},
	color.RGBA{139, 0, 0, 255},
	color.RGBA{233, 150, 122, 255},
	color.RGBA{143, 188, 143, 255},
	color.RGBA{ 72, 61, 139, 255},
	color.RGBA{ 47, 79, 79, 255},
	color.RGBA{ 47, 79, 79, 255},
	color.RGBA{ 0, 206, 209, 255},
	color.RGBA{148, 0, 211, 255},
	color.RGBA{255, 20, 147, 255},
	color.RGBA{ 0, 191, 255, 255},
	color.RGBA{105, 105, 105, 255},
	color.RGBA{ 30, 144, 255, 255},
	color.RGBA{178, 34, 34, 255},
	color.RGBA{ 34, 139, 34, 255},
	color.RGBA{255, 0, 255, 255},
	color.RGBA{255, 215, 0, 255},
	color.RGBA{218, 165, 32, 255},
	color.RGBA{ 0, 128, 0, 255},
	color.RGBA{173, 255, 47, 255},
	color.RGBA{255, 105, 180, 255},
	color.RGBA{205, 92, 92, 255},
	color.RGBA{ 75, 0, 130, 255},
	color.RGBA{240, 230, 140, 255},
	color.RGBA{124, 252, 0, 255},
	color.RGBA{173, 216, 230, 255},
	color.RGBA{240, 128, 128, 255},
	color.RGBA{144, 238, 144, 255},
	color.RGBA{255, 182, 193, 255},
	color.RGBA{255, 160, 122, 255},
	color.RGBA{ 32, 178, 170, 255},
	color.RGBA{135, 206, 250, 255},
	color.RGBA{119, 136, 153, 255},
	color.RGBA{176, 196, 222, 255},
	color.RGBA{ 0, 255, 0, 255},
	color.RGBA{ 50, 205, 50, 255},
	color.RGBA{255, 0, 255, 255},
	color.RGBA{128, 0, 0, 255},
	color.RGBA{102, 205, 170, 255},
	color.RGBA{ 0, 0, 205, 255},
	color.RGBA{186, 85, 211, 255},
	color.RGBA{147, 112, 219, 255},
	color.RGBA{ 60, 179, 113, 255},
	color.RGBA{123, 104, 238, 255},
	color.RGBA{ 0, 250, 154, 255},
	color.RGBA{ 72, 209, 204, 255},
	color.RGBA{199, 21, 133, 255},
	color.RGBA{ 25, 25, 112, 255},
	color.RGBA{128, 128, 0, 255},
	color.RGBA{107, 142, 35, 255},
	color.RGBA{255, 165, 0, 255},
	color.RGBA{255, 69, 0, 255},
	color.RGBA{218, 112, 214, 255},
	color.RGBA{238, 232, 170, 255},
	color.RGBA{152, 251, 152, 255},
	color.RGBA{175, 238, 238, 255},
	color.RGBA{219, 112, 147, 255},
	color.RGBA{255, 218, 185, 255},
	color.RGBA{205, 133, 63, 255},
	color.RGBA{255, 192, 203, 255},
	color.RGBA{221, 160, 221, 255},
	color.RGBA{176, 224, 230, 255},
	color.RGBA{128, 0, 128, 255},
	color.RGBA{255, 0, 0, 255},
	color.RGBA{188, 143, 143, 255},
	color.RGBA{ 65, 105, 225, 255},
	color.RGBA{139, 69, 19, 255},
	color.RGBA{250, 128, 114, 255},
	color.RGBA{244, 164, 96, 255},
	color.RGBA{ 46, 139, 87, 255},
	color.RGBA{160, 82, 45, 255},
	color.RGBA{192, 192, 192, 255},
	color.RGBA{135, 206, 235, 255},
	color.RGBA{106, 90, 205, 255},
	color.RGBA{112, 128, 144, 255},
	color.RGBA{ 0, 255, 127, 255},
	color.RGBA{ 70, 130, 180, 255},
	color.RGBA{210, 180, 140, 255},
	color.RGBA{ 0, 128, 128, 255},
	color.RGBA{216, 191, 216, 255},
	color.RGBA{255, 99, 71, 255},
	color.RGBA{ 64, 224, 208, 255},
	color.RGBA{238, 130, 238, 255},
	color.RGBA{245, 222, 179, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{154, 205, 50, 255},
}