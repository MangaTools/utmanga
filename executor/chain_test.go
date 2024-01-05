package executor

import (
	"errors"
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainRegular(t *testing.T) {
	colors := []color.RGBA{
		{R: 255},
		{G: 255},
		{B: 255},
	}

	testCases := []struct {
		Name           string
		PixelColors    []color.RGBA
		ExecutionFuncs []ExecutionFunc
	}{
		{
			Name:        "single",
			PixelColors: colors[0:1],
			ExecutionFuncs: []ExecutionFunc{
				createSetPixelImageColorFunc(t, colors[0], 0, 0),
			},
		},
		{
			Name:        "two",
			PixelColors: colors[0:2],
			ExecutionFuncs: []ExecutionFunc{
				createSetPixelImageColorFunc(t, colors[0], 0, 0),
				createSetPixelImageColorFunc(t, colors[1], 1, 0),
			},
		},
		{
			Name:        "three",
			PixelColors: colors[0:3],
			ExecutionFuncs: []ExecutionFunc{
				createSetPixelImageColorFunc(t, colors[0], 0, 0),
				createSetPixelImageColorFunc(t, colors[1], 1, 0),
				createSetPixelImageColorFunc(t, colors[2], 2, 0),
			},
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.Name, func(t *testing.T) {
			executeFunc := ChainExecutions(testCase.ExecutionFuncs...)
			resultImg, err := executeFunc(create3pxRgbaImage(t))
			require.NoError(t, err)

			for x, color := range testCase.PixelColors {
				imgColor := resultImg.At(x, 0)

				assert.EqualValues(t, color, imgColor)
			}

			for x := len(testCase.PixelColors); x < 3; x++ {
				imgColor := resultImg.At(x, 0)

				assert.EqualValues(t, color.RGBA{}, imgColor)
			}
		})
	}
}

func TestChainZeroExecution(t *testing.T) {
	executeFunc := ChainExecutions()

	resultImg, err := executeFunc(create3pxRgbaImage(t))
	require.NoError(t, err)

	for x := 0; x < 3; x++ {
		imgColor := resultImg.At(x, 0)

		assert.EqualValues(t, color.RGBA{}, imgColor)
	}
}

func TestChainErrorExecution(t *testing.T) {
	failExecution := ExecutionFunc(func(img image.Image) (image.Image, error) {
		return img, errors.New("expected error")
	})

	executeFunc := ChainExecutions(failExecution)

	resultImg, err := executeFunc(create3pxRgbaImage(t))
	require.Error(t, err)
	require.Nil(t, resultImg)
}

func createSetPixelImageColorFunc(t *testing.T, color color.RGBA, x, y int) ExecutionFunc {
	return func(img image.Image) (image.Image, error) {
		t.Helper()

		rgba := img.(*image.RGBA)
		rgba.SetRGBA(x, y, color)

		return rgba, nil
	}
}

func create3pxRgbaImage(t *testing.T) *image.RGBA {
	t.Helper()

	return image.NewRGBA(image.Rect(0, 0, 3, 1))
}
