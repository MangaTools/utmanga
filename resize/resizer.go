package resize

import (
	"fmt"
	"image"
	"math"

	"github.com/MangaTools/utmanga/dto"
	"github.com/disintegration/gift"
)

var modeToResampleFilter = map[string]gift.Resampling{
	dto.ResizeModeBox:             gift.BoxResampling,
	dto.ResizeModeCubic:           gift.CubicResampling,
	dto.ResizeModeLanczos:         gift.LanczosResampling,
	dto.ResizeModeLinear:          gift.LinearResampling,
	dto.ResizeModeNearestNeighbor: gift.NearestNeighborResampling,
}

type Resizer struct {
	settings dto.ResizeSettings
	filter   gift.Resampling
}

func NewResizer(settings dto.ResizeSettings) *Resizer {
	filter, ok := modeToResampleFilter[settings.Mode]
	if !ok {
		panic(fmt.Sprintf("filter %s is not in list", settings.Mode))
	}

	return &Resizer{
		settings: settings,
		filter:   filter,
	}
}

func (r *Resizer) Execute(img image.Image) (image.Image, error) {
	bounds := img.Bounds()

	targetHeight := r.getTaretHeight(bounds)
	stepValue := (targetHeight - bounds.Dy()) / r.settings.Steps

	g := gift.New()

	// NOTE(shadream): evaluate step - 1 resize and then resize to target value outside of loop to more accurate final size.
	for i := 1; i < r.settings.Steps; i++ {
		newHeight := bounds.Dy() + (stepValue * i)
		g.Add(gift.Resize(0, newHeight, r.filter))
	}

	if r.settings.TargetWidth != 0 {
		g.Add(gift.Resize(r.settings.TargetWidth, 0, r.filter))
	} else {
		g.Add(gift.Resize(0, targetHeight, r.filter))
	}

	g.Add(gift.Grayscale())

	resultImage := image.NewGray(g.Bounds(img.Bounds()))

	g.Draw(resultImage, img)

	return img, nil
}

func (r *Resizer) getTaretHeight(bounds image.Rectangle) int {
	var targetHeight int
	switch {
	case r.settings.Coefficient < 0:
		targetHeight = bounds.Dy() / (-r.settings.Coefficient)
	case r.settings.Coefficient > 0:
		targetHeight = bounds.Dy() * r.settings.Coefficient
	case r.settings.TargetHeigth != 0:
		targetHeight = r.settings.TargetHeigth
	case r.settings.TargetWidth != 0:
		coefficient := float64(r.settings.TargetWidth) / float64(bounds.Dx())
		targetHeight = int(math.Round(float64(bounds.Dx()) * coefficient))
	}

	return targetHeight
}
