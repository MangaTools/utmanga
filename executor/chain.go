package executor

import (
	"fmt"
	"image"
)

// ChainExecutions create ExecutionFunc that execute all executions one by one.
// If one of executions returns error result ExecutionFunc stops execution and return error
func ChainExecutions(executions ...ExecutionFunc) ExecutionFunc {
	return func(image image.Image) (image.Image, error) {
		currentImage := image
		for _, execute := range executions {
			newImage, err := execute(currentImage)
			if err != nil {
				return nil, fmt.Errorf("one of the steps in execution failed: %w", err)
			}

			currentImage = newImage
		}

		return currentImage, nil
	}
}
