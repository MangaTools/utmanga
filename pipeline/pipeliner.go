package pipeline

import (
	"image"

	"github.com/MangaTools/utmanga/dto"
	"github.com/MangaTools/utmanga/executor"
	"github.com/MangaTools/utmanga/level"
	"github.com/MangaTools/utmanga/resize"
	"github.com/MangaTools/utmanga/screentone"
)

type Pipeliner struct {
	execution executor.ExecutionFunc
}

func NewPipeliner(settings dto.PipelineFileSettings) *Pipeliner {
	return &Pipeliner{
		execution: executor.ChainExecutions(getExecutions(settings)...),
	}
}

func (p *Pipeliner) Execute(img image.Image) (image.Image, error) {
	return p.execution(img)
}

func getExecutions(settings dto.PipelineFileSettings) []executor.ExecutionFunc {
	var executions []executor.ExecutionFunc
	for _, item := range settings.Items {
		switch item.Type {
		case dto.PipelineTypeLevel:
			levelSettings := item.Settings.(dto.LevelSettings)
			leveler := level.NewLeveler(levelSettings)

			executions = append(executions, leveler.Execute)
		case dto.PipelineTypeResize:
			resizeSettings := item.Settings.(dto.ResizeSettings)
			resizer := resize.NewResizer(resizeSettings)
			executions = append(executions, resizer.Execute)

		case dto.PipelineTypeScrentone:
			screentoneSettings := item.Settings.(dto.ScreentoneSettings)
			screentoner := screentone.NewScreentoner(screentoneSettings)
			executions = append(executions, screentoner.Execute)
		}
	}

	return executions
}
