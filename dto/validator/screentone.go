package validator

import (
	"github.com/MangaTools/utmanga/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateScreentone(settings dto.ScreentoneSettings) error {
	return validation.ValidateStruct(&settings,
		validation.Field(&settings.DotSize, validation.Min(2), validation.Max(100)),
	)
}
