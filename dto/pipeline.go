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
	var items []pipelineItem
	if err := json.Unmarshal(b, &items); err != nil {
		return fmt.Errorf("parse PipelineFileSetting array: %w", err)
	}

	var resultItems []PipelineItem

	for _, item := range items {
		var pipelineItem *PipelineItem
		var err error

		switch item.Type {
		case PipelineTypeLevel:
			var levelSettings LevelSettings
			pipelineItem, err = createPipelineItem(item.Settings, &levelSettings, PipelineTypeLevel)
		case PipelineTypeResize:
			var resizeSettings ResizeSettings
			pipelineItem, err = createPipelineItem(item.Settings, &resizeSettings, PipelineTypeResize)
		case PipelineTypeScrentone:
			var screentoneSettings ScreentoneSettings
			pipelineItem, err = createPipelineItem(item.Settings, &screentoneSettings, PipelineTypeScrentone)
		default:
			return errors.New("undefined item.Type")
		}
		if err != nil {
			return err
		}

		resultItems = append(resultItems, *pipelineItem)
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
	Type     string
	Settings json.RawMessage
}

type PipelineItem struct {
	Type     string           `json:"type"`
	Settings PipelineItemData `json:"settings"`
}

type PipelineItemData interface{}
