package resize

import (
	"fmt"
	"image"
	"math"

	"github.com/MangaTools/utmanga/dto"
	"github.com/disintegration/imaging"
)

var modeToResampleFilter = map[string]imaging.ResampleFilter{
	dto.ResizeModeBox:               imaging.Box,
	dto.ResizeModeCatmullRom:        imaging.CatmullRom,
	dto.ResizeModeLanczos:           imaging.Lanczos,
	dto.ResizeModeLinear:            imaging.Linear,
	dto.ResizeModeMitchellNetravali: imaging.MitchellNetravali,
	dto.ResizeModeNearestNeighbor:   imaging.NearestNeighbor,
}

type Resizer struct {
	settings dto.ResizeSettings
	filter   imaging.ResampleFilter
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

	// NOTE(shadream): evaluate step - 1 resize and then resize to target value outside of loop to more accurate final size.
	for i := 1; i < r.settings.Steps; i++ {
		newHeight := bounds.Dy() + (stepValue * i)
		img = imaging.Resize(img, 0, newHeight, r.filter)
	}

	if r.settings.TargetWidth != 0 {
		img = imaging.Resize(img, r.settings.TargetWidth, 0, r.filter)
	} else {
		img = imaging.Resize(img, 0, targetHeight, r.filter)
	}

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
