package ffmpeg

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func getLayers(scriptsPath, inputVideo, codecType string) ([]Stream, error) {
	var (
		layerOutput         = LayerOutput{Streams: make([]Stream, 0)}
		extractLayersScript = fmt.Sprintf("%s%s", scriptsPath, "/get_layers.sh")

		index = 0
	)

	cmd, err := exec.Command("/bin/sh", extractLayersScript, inputVideo).Output()
	if err != nil {
		return layerOutput.Streams, err
	}

	err = json.Unmarshal(cmd, &layerOutput)
	if err != nil {
		return layerOutput.Streams, err
	}

	for _, stream := range layerOutput.Streams {
		if stream.CodecType == codecType {
			layerOutput.Streams[index] = stream
			index++
		}
	}

	return layerOutput.Streams[:index], nil
}
