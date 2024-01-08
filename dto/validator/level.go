package validator

import (
	"github.com/MangaTools/utmanga/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateLevelSettings(settings dto.LevelSettings) error {
	return validation.ValidateStruct(&settings,
		validation.Field(&settings.LowInput, validation.Min(0), validation.Max(settings.HighInput)),
		validation.Field(&settings.HighInput, validation.Min(settings.LowInput), validation.Max(255)),
		validation.Field(&settings.Gamma, validation.Min(float32(0.01)), validation.Max(float32(9.99))),
		validation.Field(&settings.LowOutput, validation.Min(0), validation.Max(settings.HighOutput)),
		validation.Field(&settings.HighOutput, validation.Min(settings.LowOutput), validation.Max(255)),
		validation.Field(&settings.DiapasonBlack, validation.Min(-1), validation.Max(255-settings.DiapasonWhite)),
		validation.Field(&settings.DiapasonWhite, validation.Min(-1), validation.Max(255-settings.DiapasonBlack)),
	)
}
