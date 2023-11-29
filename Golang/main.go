package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/danhilltech/goyolov5"
)

func main() {
	yolov5, err := goyolov5.NewYoloV5("yolov5n.torchscript.gpu.batch1.pt", goyolov5.DeviceCPU, 640, false)
	if err != nil {
		panic(err)
	}

	f, err := os.Open("path/to/my/image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	tensor := goyolov5.NewTensorFromImage(input)

	outTensor := goyolov5.NewTensorFromImage(tensor)

	predictions, err := yolov5.Infer(tensor, 0.5, 0.4, outTensor)
	if err != nil {
		panic(err)
	}

	fmt.Println(predictions)

	f2, err := os.Create("path/to/my/annotated_image.png")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	if err = png.Encode(f2, outTensor); err != nil {
		panic(err)
	}
}
