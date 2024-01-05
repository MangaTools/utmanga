package executor

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const recursiveRootFolder = "./testdata/recursive"

func TestRegular(t *testing.T) {
	files, err := getImagesPaths(recursiveRootFolder, acceptedImageExtensions, false)
	require.NoError(t, err)

	assert.Len(t, files, 3)

	files, err = getImagesPaths(recursiveRootFolder, acceptedImageExtensions, true)
	require.NoError(t, err)
	assert.Len(t, files, 12)

	files, err = getImagesPaths(filepath.Join(recursiveRootFolder, "not_exists"), acceptedImageExtensions, true)
	assert.Error(t, err)
}
