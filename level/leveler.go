package level

import (
	"image"
	"image/color"
	"math"

	"github.com/MangaTools/utmanga/dto"
	"github.com/disintegration/gift"
)

const (
	diapasonBlackBlur          = float32(1)
	diapasonBlackBlurThreshold = uint8(230)

	diapasonWhiteBlur = 3
)

type Leveler struct {
	settings dto.LevelSettings
}

func NewLeveler(settings dto.LevelSettings) *Leveler {
	return &Leveler{
		settings: settings,
	}
}

func (l *Leveler) Execute(img image.Image) (image.Image, error) {
	grayscaleFilter := gift.New(gift.Grayscale())
	grayImage := image.NewGray(grayscaleFilter.Bounds(img.Bounds()))
	grayscaleFilter.Draw(grayImage, img)

	resultImage := image.NewGray(img.Bounds())
	grayPixel := color.Gray{}
	gammaCorrection := 1 / l.settings.Gamma

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			imgGrayPixel := int(grayImage.GrayAt(x, y).Y)

			value := float64(imgGrayPixel-l.settings.LowInput) / float64(l.settings.HighInput-l.settings.LowInput)

			if l.settings.Gamma != 1 {
				value = math.Pow(value, gammaCorrection)
			}

			value = value*(float64(l.settings.HighOutput)-float64(l.settings.LowOutput)) + float64(l.settings.LowOutput)

			grayPixel.Y = uint8(min(255, max(0, value)))

			resultImage.SetGray(x, y, grayPixel)
		}
	}

	l.executeDiapasonBlack(resultImage)
	l.executeDiapasonWhite(resultImage, grayImage)

	return resultImage, nil
}

func (l *Leveler) executeDiapasonBlack(grayImg *image.Gray) {
	if l.settings.DiapasonBlack == 0 {
		return
	}

	blackThreshold := uint8(l.settings.DiapasonBlack)
	blackMask := image.NewGray(grayImg.Bounds())
	WhitePixel := color.Gray{Y: 255}
	for y := 0; y < grayImg.Bounds().Dy(); y++ {
		for x := 0; x < grayImg.Bounds().Dx(); x++ {
			imgColor := grayImg.GrayAt(x, y)

			if imgColor.Y > blackThreshold {
				blackMask.SetGray(x, y, WhitePixel)
			}
		}
	}

	g := gift.New(gift.GaussianBlur(diapasonBlackBlur))

	blurredMask := image.NewGray(g.Bounds(blackMask.Rect))
	g.Draw(blurredMask, blackMask)

	blackPixel := color.Gray{Y: 0}

	for y := 0; y < blurredMask.Bounds().Dy(); y++ {
		for x := 0; x < blurredMask.Bounds().Dx(); x++ {
			value := blurredMask.GrayAt(x, y).Y
			if value < diapasonBlackBlurThreshold {
				grayImg.SetGray(x, y, blackPixel)
			}
		}
	}
}

func (l *Leveler) executeDiapasonWhite(resultImage *image.Gray, sourceGrayImage *image.Gray) {
	if l.settings.DiapasonWhite == 0 {
		return
	}

	minDiapason := uint8(255 - l.settings.DiapasonWhite)

	medianGrayFilter := gift.New(gift.Median(diapasonWhiteBlur, false))
	medianGrayImage := image.NewGray(medianGrayFilter.Bounds(sourceGrayImage.Rect))
	medianGrayFilter.Draw(medianGrayImage, sourceGrayImage)

	whitePixel := color.Gray{Y: 255}

	for y := 0; y < resultImage.Bounds().Dy(); y++ {
		for x := 0; x < resultImage.Bounds().Dx(); x++ {
			medianGray := medianGrayImage.GrayAt(x, y).Y

			if medianGray > minDiapason {
				resultImage.SetGray(x, y, whitePixel)
			}
		}
	}
}
