package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDot(t *testing.T) {
	testCases := []struct {
		Name          string
		Size          int
		Inverted      bool
		ExpectedSlice []byte
	}{
		{
			Name:     "regular 2",
			Size:     2,
			Inverted: false,
			ExpectedSlice: []byte{
				129, 171,
				213, 255,
			},
		},
		{
			Name:     "regular 3",
			Size:     3,
			Inverted: false,
			ExpectedSlice: []byte{
				129, 192, 145,
				208, 255, 224,
				161, 239, 176,
			},
		},
		{
			Name:     "inverted 2",
			Size:     2,
			Inverted: true,
			ExpectedSlice: []byte{
				128, 86,
				43, 1,
			},
		},
		{
			Name:     "inverted 3",
			Size:     3,
			Inverted: true,
			ExpectedSlice: []byte{
				128, 65, 112,
				49, 1, 33,
				96, 17, 80,
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			slice := generateDot(tc.Size, tc.Inverted)
			assert.EqualValues(t, slice, tc.ExpectedSlice)
		})
	}
}
