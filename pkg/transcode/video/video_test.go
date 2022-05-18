package video_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoDUrationGet(t *testing.T) {
	mainPath := "/home/samandar/Downloads/videos"

	files, err := os.ReadDir(mainPath)
	assert.NoError(t, err)

	for _, file := range files {
		fmt.Println(filepath.Join(mainPath, file.Name()))
		_, err := videoTest.GetVideoDuration(filepath.Join(mainPath, file.Name()))
		assert.NoError(t, err)
	}
}
