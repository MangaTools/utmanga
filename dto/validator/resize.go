package validator

import (
	"github.com/MangaTools/utmanga/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateResizeSettings(settings dto.ResizeSettings) error {
	return validation.ValidateStruct(&settings,
		validation.Field(&settings.Steps, validation.Min(1), validation.Max(10)),
		validation.Field(&settings.Mode, validation.In(dto.ResizeModeLanczos, dto.ResizeModeCubic,
			dto.ResizeModeLinear, dto.ResizeModeBox, dto.ResizeModeNearestNeighbor), validation.Required),
		validation.Field(&settings.Coefficient, validation.In(-4, -2, 0, 2, 4)),
		validation.Field(&settings.TargetWidth, validation.Min(0)),
		validation.Field(&settings.TargetHeigth, validation.Min(0)),
	)
}
