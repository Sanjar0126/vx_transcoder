package ffmpeg

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func getLayers(scriptsPath, inputVideo, layerType string) ([]Streams, error) {
	var (
		layerOutput = LayerOutput{Streams: make([]Streams, 0)}

		extractLayersScript = fmt.Sprintf("%s%s", scriptsPath, "/get_layers.sh")
	)

	cmd, err := exec.Command("/bin/sh", extractLayersScript, inputVideo).Output()
	if err != nil {
		return layerOutput.Streams, err
	}

	err = json.Unmarshal(cmd, &layerOutput)
	if err != nil {
		return layerOutput.Streams, err
	}

	return layerOutput.Streams, nil
}
