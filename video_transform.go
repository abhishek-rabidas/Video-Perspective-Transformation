package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"math"
)

func main() {

	frames, err := gocv.VideoCaptureFile("./asset/vid.mp4")

	if err != nil {
		panic(err)
	}

	window := gocv.NewWindow("Transposed")
	defer window.Close()

	mat := gocv.NewMat()
	defer mat.Close()

	origImg := []image.Point{
		{377, 53},   // top-left
		{35, 679},   // bottom-left
		{1863, 677}, // bottom-right
		{1601, 53},  // top-right
	}

	fmt.Println("Lane width: ", distance(origImg[0], origImg[3]))
	fmt.Println("Lane length: ", distance(origImg[0], origImg[1]))

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

	originalWindow := gocv.NewWindow("original")

	defer originalWindow.Close()

	for {
		frames.Read(&mat)
		transform := gocv.GetPerspectiveTransform(gocv.NewPointVectorFromPoints(origImg), newImg)
		perspective := gocv.NewMat()
		gocv.WarpPerspective(mat, &perspective, transform, image.Point{width, height})

		originalWindow.IMShow(mat)
		window.IMShow(perspective)

		window.WaitKey(1)
	}
}

func distance(p1, p2 image.Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}
