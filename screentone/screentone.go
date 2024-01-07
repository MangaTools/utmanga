package screentone

import (
	"image"
	"image/color"

	"github.com/MangaTools/utmanga/dto"
	"github.com/MangaTools/utmanga/screentone/algorithm"
)

type Screentoner struct {
	matrix []byte
	size   int
}

func NewScreentoner(settings dto.ScreentoneSettings) *Screentoner {
	matrix, size := algorithm.GenerateDotScreentone(settings.DotSize)
	return &Screentoner{
		matrix: matrix,
		size:   size,
	}
}

func (s *Screentoner) Execute(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	resultImage := image.NewGray(bounds)
	grayConverter := color.GrayModel

	grayPixel := color.Gray{}

	for y := 0; y < bounds.Dy(); y++ {
		yIndex := (y % s.size) * s.size
		for x := 0; x < bounds.Dx(); x++ {
			grayColor := grayConverter.Convert(img.At(x, y)).(color.Gray).Y

			if grayColor < s.matrix[(x%s.size)+yIndex] {
				grayPixel.Y = 0
			} else {
				grayPixel.Y = 255
			}

			resultImage.SetGray(x, y, grayPixel)
		}
	}

	return resultImage, nil
}
