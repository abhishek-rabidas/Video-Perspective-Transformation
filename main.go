package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
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
		{352, 32},  // top-left
		{265, 294}, // bottom-left
		{660, 272}, // bottom-right
		{468, 24},  // top-right
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

	for {

		video.Read(&mat)

		for _, point := range origImg {
			createCircles(&mat, point)
		}

		window.IMShow(mat)
		transform := gocv.GetPerspectiveTransform(gocv.NewPointVectorFromPoints(origImg), newImg)
		perspective := gocv.NewMat()
		gocv.WarpPerspective(mat, &perspective, transform, image.Point{width, height})
		tranformed_window.IMShow(perspective)

		window.WaitKey(1)
	}

}

func createCircles(mat *gocv.Mat, point image.Point) {
	gocv.Circle(mat, point, 2, color.RGBA{R: 255}, 2)
}
