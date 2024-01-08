package validator

import (
	"errors"

	"github.com/MangaTools/utmanga/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateResizeSettings(settings dto.ResizeSettings) error {
	err := validation.ValidateStruct(&settings,
		validation.Field(&settings.Steps, validation.Min(1), validation.Max(10)),
		validation.Field(&settings.Mode, validation.In(dto.ResizeModeLanczos, dto.ResizeModeCubic,
			dto.ResizeModeLinear, dto.ResizeModeBox, dto.ResizeModeNearestNeighbor), validation.Required),
		validation.Field(&settings.Coefficient, validation.In(-4, -2, 0, 2, 4)),
		validation.Field(&settings.TargetWidth, validation.Min(0)),
		validation.Field(&settings.TargetHeigth, validation.Min(0)),
	)
	if err != nil {
		return err
	}

	var settedFilds int
	if settings.Coefficient != 0 {
		settedFilds++
	}
	if settings.TargetHeigth != 0 {
		settedFilds++
	}
	if settings.TargetWidth != 0 {
		settedFilds++
	}

	if settedFilds == 0 {
		return errors.New("one item within coefficient, target_heigth or target_width should not be default")
	}

	if settedFilds > 1 {
		return errors.New("only one item within coefficient, target_heigth or target_width should be set")
	}

	return nil
}
