package barcode

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type QRCode struct {
	// Version      int      // <-- This represents the QR code version
	// ErrorCorrectionLevel string
	// Data         string
	// Modules      [][]bool
}

func (qr *QRCode) GenQrCode() {
	// fixed  code
	qrPattern := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1},
	}
	quietZone := 4
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	patternSize := len(qrPattern)
	fullSize := patternSize + 2*quietZone
	img := image.NewRGBA(image.Rect(0, 0, fullSize, fullSize))
	for y := 0; y < fullSize; y++ {
		for x := 0; x < fullSize; x++ {
			img.Set(x, y, white)
		}
	}

	for y, row := range qrPattern {
		for x, col := range row {
			if col == 1 {
				img.Set(x+quietZone, y+quietZone, black)
			}
		}

	}
	qr.saveFile(img)
}

func (*QRCode) saveFile(img *image.RGBA) {
	file, err := os.Create("./qr.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func (qr *QRCode) Decode() [][]int {
	fileName := "qr.png"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	qrPattern := make([][]int, height)
	for i := range qrPattern {
		qrPattern[i] = make([]int, width)
	}
	threshold := uint32(128 * 128 * 128)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Simple grayscale conversion and thresholding
			luminance := (r + g + b) / 3
			if luminance < threshold {
				qrPattern[y][x] = 1 // Black
			} else {
				qrPattern[y][x] = 0 // White
			}
		}
	}
	return qrPattern
}

func NewQrCode() QRCode {
	return QRCode{}
}
