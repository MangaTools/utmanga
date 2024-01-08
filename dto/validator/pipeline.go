package validator

import (
	"errors"
	"fmt"

	"github.com/MangaTools/utmanga/dto"
)

func ValidatePipelineSettings(settings dto.PipelineFileSettings) error {
	for _, item := range settings.Items {
		switch item.Type {
		case dto.PipelineTypeLevel:
			levelSettings := item.Settings.(dto.LevelSettings)
			if err := ValidateLevelSettings(levelSettings); err != nil {
				return err
			}
		case dto.PipelineTypeResize:
			resizeSettings := item.Settings.(dto.ResizeSettings)
			if err := ValidateResizeSettings(resizeSettings); err != nil {
				return err
			}
		case dto.PipelineTypeScrentone:
			screentoneSettings := item.Settings.(dto.ScreentoneSettings)
			if err := ValidateScreentoneSettings(screentoneSettings); err != nil {
				return err
			}
		default:
			return errors.New(fmt.Sprintf("can not validate type %s", item.Type))
		}
	}

	return nil
}
