package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

var lineHeight = 1
var lineWidthMax = 75
var lineWidthMin = 10

func main() {
	startTime := time.Now()
	rand.Seed(startTime.UnixNano())

	img, err := gg.LoadImage("source.jpg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	imageSize := img.Bounds().Size()
	rows := (imageSize.Y / lineHeight)
	ctx := gg.NewContext(imageSize.X, rows*lineHeight)

	ctx.DrawRectangle(0, 0, float64(imageSize.X), float64(imageSize.Y))
	ctx.SetRGB255(255, 255, 255)
	ctx.Fill()

	for y := 0; y < rows*lineHeight; y += lineHeight {
		for x := 0; x < imageSize.X-1; {
			lineLength := rand.Intn(lineWidthMax-lineWidthMin+1) + lineWidthMin
			if x+lineLength >= imageSize.X {
				lineLength = imageSize.X - x
			}

			ctx.DrawRectangle(float64(x), float64(y), float64(lineLength), float64(lineHeight))
			ctx.SetColor(getColor(img, x, y, lineLength, lineHeight))
			ctx.Fill()

			x += lineLength
		}
	}

	ctx.SavePNG("result.png")
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime.Nanoseconds()/1000000, "ms")
}

func getColor(img image.Image, x, y, width, height int) color.Color {
	// c := img.At(x, y)

	var colors = make([]color.Color, width*height)
	index := 0
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			colors[index] = img.At(x+i, y+j)
			index++
		}
	}

	c := getAverageColor(colors)

	return c
}

func getAverageColor(colors []color.Color) color.Color {
	var cl color.Color

	sr, sg, sb, _ := colors[0].RGBA()
	// sl := getBrightness(toRGBChannel(sr), toRGBChannel(sg), toRGBChannel(sb))
	er, eg, eb, _ := colors[len(colors)-1].RGBA()
	// el := getBrightness(toRGBChannel(er), toRGBChannel(eg), toRGBChannel(eb))

	cl = color.NRGBA{
		R: toRGBChannel((sr + er) / 2),
		G: toRGBChannel((sg + eg) / 2),
		B: toRGBChannel((sb + eb) / 2),
		// A: toRGBChannel((sa + ea) / 2),
		// A: ((sl + el) / 2),
		A: 255,
	}

	return cl
}

func toRGBChannel(n uint32) uint8 {
	return uint8(n / 0x101)
}

func getBrightness(r, g, b uint8) uint8 {
	return 245 + uint8((1-((float32(r)*float32(g)*float32(b))/16581375)))*10
}
