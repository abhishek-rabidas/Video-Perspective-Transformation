package main

import (
	"gocv.io/x/gocv"
	"image"
	"math"
)

func main() {

	img := gocv.IMRead("./asset/road.png", gocv.IMReadAnyColor)
	//Line
	/*	start := image.Point{447, 1017}
		end := image.Point{444, 1473}

		gocv.Line(&img, start, end, color.RGBA{255, 0, 0, 0}, 2)*/

	origImg := []image.Point{
		{377, 53},   // top-left
		{35, 679},   // bottom-left
		{1863, 677}, // bottom-right
		{1601, 53},  // top-right
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

	// Draw the line on the image

	//fmt.Println("Distance: ", distance(start, end))

	/*	gocv.Line(&perspective, image.Point{
			X: 0,
			Y: 0,
		}, image.Point{
			X: 431,
			Y: 0,
		}, color.RGBA{0, 255, 0, 0}, 2)
		fmt.Println("Distance: ", distance(image.Point{
			X: 0,
			Y: 0,
		}, image.Point{
			X: 431,
			Y: 0,
		}))*/
	gocv.IMWrite("transformed.jpg", perspective)

}

func distance(p1, p2 image.Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}
