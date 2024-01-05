package executor

import (
	"fmt"
	"image"
	"image/png"
	"io"

	"github.com/samber/lo"
)

type EncoderFunc func(w io.Writer, img image.Image) error

const (
	pngEncoderName = "png"
)

var encoderList = map[string]bool{
	pngEncoderName: true,
}

func mustGetEncoderByString(encoderType string) EncoderFunc {
	switch encoderType {
	case pngEncoderName:
		return png.Encode
	default:
		panic(fmt.Sprintf("%s encoder type is not exist", encoderType))
	}
}

func IsValidEncoderType(encodeType string) bool {
	_, ok := encoderList[encodeType]

	return ok
}

func GetListEncoderTypes() []string {
	return lo.Keys(encoderList)
}

var DefaultEncoder EncoderFunc = png.Encode
