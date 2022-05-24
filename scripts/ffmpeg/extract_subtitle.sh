#!/bin/bash
input_video=$1
track_number=$3
segment_time=$4
output_path=$5

echo $input_video
echo $lang
echo $track_number
echo $segment_time
echo $output_path

output=$output_path/subtitles/$lang/subtitle.vtt
segment_output=$output_path/subtitles/$lang

ffmpeg -y -i $input_video -map 0:s:$track_number -c:s webvtt $output

ffmpeg \
    -i $output -f segment -segment_time $segment_time -segment_list_size 0 \
    -segment_list $segment_output/sub.m3u8 \
    -segment_format webvtt -scodec copy $segment_output/sub-%d.vtt
