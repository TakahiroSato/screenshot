package main

import (
	"flag"
	"image"
	"image/png"
	"os"
	"strconv"

	"github.com/TakahiroSato/imgconv"
	"github.com/kbinani/screenshot"
)

var binary = flag.Bool("binary", false, "二値化するかどうか(default=false)")
var reverse = flag.Bool("reverse", false, "二値化白黒反転するかどうか(default=false)")
var th int

func init() {
	flag.IntVar(&th, "th", 127, "二値化しきい値(0 ~ 255)")
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 4 {
		os.Exit(1)
	}
	sx, _ := strconv.Atoi(flag.Args()[0])
	sy, _ := strconv.Atoi(flag.Args()[1])
	dx, _ := strconv.Atoi(flag.Args()[2])
	dy, _ := strconv.Atoi(flag.Args()[3])

	rect := image.Rect(sx, sy, dx, dy)
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}
	var fileName string
	if len(flag.Args()) < 5 {
		fileName = "screenshot.png"
	} else {
		fileName = flag.Args()[4]
	}

	// 二値化
	if *binary {
		imgconv.ToBinary(img, uint8(th), *reverse).SaveAsPng(fileName)
	} else {
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
	}

}
