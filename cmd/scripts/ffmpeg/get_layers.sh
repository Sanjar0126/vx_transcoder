#!/bin/bash
input_video=$1
ffprobe -v fatal -show_streams "$input_video" -print_format json
