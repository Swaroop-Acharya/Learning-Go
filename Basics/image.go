package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"fmt"
)

func decode() {
	f, _ := os.Open("output.png")
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Image bounds:", img.Bounds())
	fmt.Println("Color at (10,10):", img.At(10, 10))
}

func createImg(){
	// Create a 100x100 image
	img := image.NewRGBA(image.Rect(50, 50, 100, 100))

	// Fill it with red color
	red := color.RGBA{255, 0, 0, 255}
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, red)
		}
	}

	// Save it as output.png
	f, _ := os.Create("output.png")
	defer f.Close()
	png.Encode(f, img)


}

func main() {
	createImg()
	decode()
}
