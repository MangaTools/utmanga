package validator

import (
	"runtime"

	"github.com/MangaTools/utmanga/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateExecutorSettings(settings dto.ExecutorSettings) error {
	return validation.ValidateStruct(&settings,
		validation.Field(&settings.MaxGoroutines, validation.Min(1), validation.Max(runtime.NumCPU())))
}
