package ffmpeg

import (
	"fmt"
	"os/exec"
)

type FolderOpts struct {
	OutputPath   string
	AudioList    string
	SubtitleList string
	VideoList    string
	ScriptPath   string
}

func CreateFolders(opts FolderOpts) error {
	var (
		err        error
		scriptPath = fmt.Sprintf("%s%s", opts.ScriptPath, "/ffmpeg/gen_folder.sh")
	)

	_, err = exec.Command("/bin/bash", scriptPath, opts.OutputPath, opts.AudioList,
		opts.VideoList, opts.SubtitleList).Output()

	return err
}
