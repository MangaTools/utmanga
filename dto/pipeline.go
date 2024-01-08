package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	PipelineTypeLevel     = "level"
	PipelineTypeResize    = "resize"
	PipelineTypeScrentone = "screentone"
)

type PipelineFileSettings struct {
	Items []PipelineItem `json:"items"`
}

func (p *PipelineFileSettings) UnmarshalJSON(b []byte) error {
	settings := struct {
		Items []pipelineItem `json:"items"`
	}{}
	if err := json.Unmarshal(b, &settings); err != nil {
		return fmt.Errorf("parse PipelineFileSetting array: %w", err)
	}

	var resultItems []PipelineItem
	for _, item := range settings.Items {
		var pipelineItem PipelineItem
		var err error

		switch item.Type {
		case PipelineTypeLevel:
			var levelSettings LevelSettings
			if err := json.Unmarshal(item.Settings, &levelSettings); err != nil {
				return fmt.Errorf("parse ScreentoneSettings: %w", err)
			}
			pipelineItem = PipelineItem{
				Type:     PipelineTypeLevel,
				Settings: levelSettings,
			}
		case PipelineTypeResize:
			var resizeSettings ResizeSettings
			if err := json.Unmarshal(item.Settings, &resizeSettings); err != nil {
				return fmt.Errorf("parse ScreentoneSettings: %w", err)
			}
			pipelineItem = PipelineItem{
				Type:     PipelineTypeResize,
				Settings: resizeSettings,
			}
		case PipelineTypeScrentone:
			var screentoneSettings ScreentoneSettings
			if err := json.Unmarshal(item.Settings, &screentoneSettings); err != nil {
				return fmt.Errorf("parse ScreentoneSettings: %w", err)
			}
			pipelineItem = PipelineItem{
				Type:     PipelineTypeScrentone,
				Settings: screentoneSettings,
			}
		default:
			return errors.New("undefined item.Type")
		}
		if err != nil {
			return err
		}

		resultItems = append(resultItems, pipelineItem)
	}

	p.Items = resultItems

	return nil
}

func createPipelineItem(data json.RawMessage, marshallStruct any, pipelineType string) (*PipelineItem, error) {
	if err := json.Unmarshal(data, &marshallStruct); err != nil {
		return nil, fmt.Errorf("parse ScreentoneSettings: %w", err)
	}
	return &PipelineItem{
		Type:     PipelineTypeLevel,
		Settings: marshallStruct,
	}, nil
}

type pipelineItem struct {
	Type     string          `json:"type"`
	Settings json.RawMessage `json:"settings"`
}

type PipelineItem struct {
	Type     string           `json:"type"`
	Settings PipelineItemData `json:"settings"`
}

type PipelineItemData interface{}
