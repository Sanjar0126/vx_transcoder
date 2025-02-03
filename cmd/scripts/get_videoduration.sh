#!/bin/bash
input_video=$1

echo $input_video
ffprobe -i $input_video -show_entries format=duration -v quiet -of csv="p=0"