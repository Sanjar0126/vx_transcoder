#!/bin/bash
input_video=$1
ffprobe -v error -show_streams "$input_video" -print_format json
