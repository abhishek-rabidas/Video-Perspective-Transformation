package main

import (
	"gocv.io/x/gocv"
	"image"
	"math"
)

func main() {

	video, err := gocv.VideoCaptureFile("./asset/traffic.mp4")

	if err != nil {
		panic(err)
	}

	mat := gocv.NewMat()
	window := gocv.NewWindow("ORIGINAL")

	//transformation
	tranformed_window := gocv.NewWindow("Transformed")

	origImg := []image.Point{
		image.Point{128, 165}, // top-left
		image.Point{215, 275}, // bottom-left
		image.Point{385, 128}, // bottom-right
		image.Point{300, 40},  // top-right
	}

	heightA := math.Sqrt(math.Pow(float64(origImg[0].X-origImg[1].X), 2) + math.Pow(float64(origImg[0].Y-origImg[1].Y), 2))
	heightB := math.Sqrt(math.Pow(float64(origImg[3].X-origImg[2].X), 2) + math.Pow(float64(origImg[3].Y-origImg[2].Y), 2))
	height := int(math.Max(heightA, heightB))

	widthA := math.Sqrt(math.Pow(float64(origImg[0].X-origImg[3].X), 2) + math.Pow(float64(origImg[0].Y-origImg[3].Y), 2))
	widthB := math.Sqrt(math.Pow(float64(origImg[1].X-origImg[2].X), 2) + math.Pow(float64(origImg[1].Y-origImg[2].Y), 2))
	width := int(math.Max(widthA, widthB))

	newImg := gocv.NewPointVectorFromPoints([]image.Point{
		image.Point{0, 0},
		image.Point{0, height},
		image.Point{width, height},
		image.Point{width, 0},
	})

	for {
		video.Read(&mat)
		window.IMShow(mat)
		transform := gocv.GetPerspectiveTransform(gocv.NewPointVectorFromPoints(origImg), newImg)
		perspective := gocv.NewMat()
		gocv.WarpPerspective(mat, &perspective, transform, image.Point{width, height})
		tranformed_window.IMShow(perspective)

		window.WaitKey(1)
	}

}
