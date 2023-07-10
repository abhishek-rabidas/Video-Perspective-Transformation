package main

import (
	"gocv.io/x/gocv"
	"image"
	"math"
)

func main() {

	img := gocv.IMRead("./asset/pic.jpg", gocv.IMReadAnyColor)

	origImg := []image.Point{
		{539, 917},  // top-left
		{221, 1237}, // bottom-left
		{517, 1551}, // bottom-right
		{821, 1143}, // top-right
	}

	heightA := math.Sqrt(math.Pow(float64(origImg[0].X-origImg[1].X), 2) + math.Pow(float64(origImg[0].Y-origImg[1].Y), 2))
	heightB := math.Sqrt(math.Pow(float64(origImg[3].X-origImg[2].X), 2) + math.Pow(float64(origImg[3].Y-origImg[2].Y), 2))
	height := int(math.Max(heightA, heightB))

	widthA := math.Sqrt(math.Pow(float64(origImg[0].X-origImg[3].X), 2) + math.Pow(float64(origImg[0].Y-origImg[3].Y), 2))
	widthB := math.Sqrt(math.Pow(float64(origImg[1].X-origImg[2].X), 2) + math.Pow(float64(origImg[1].Y-origImg[2].Y), 2))
	width := int(math.Max(widthA, widthB))

	newImg := gocv.NewPointVectorFromPoints([]image.Point{
		{0, 0},
		{0, height},
		{width, height},
		{width, 0},
	})

	transform := gocv.GetPerspectiveTransform(gocv.NewPointVectorFromPoints(origImg), newImg)
	perspective := gocv.NewMat()
	gocv.WarpPerspective(img, &perspective, transform, image.Point{width, height})
	gocv.IMWrite("transformed.jpg", perspective)

}
