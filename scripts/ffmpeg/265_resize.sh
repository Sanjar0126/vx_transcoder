#!/bin/bash
input_video=$1
resize_width_height=$2
video_bit_rate=$3
hls_time=$4
output_path=$5

echo ffmpeg -y -vsync 0 -c:v hevc_cuvid -i $input_video -vf scale=$resize_width_height  \
    -map 0:v:0 -c:v libsvtav1 -b:v $video_bit_rate -g 30 -strict -2 \
    -color_primaries 1 -color_trc 1 -colorspace 1 \
    -hls_time $hls_time -hls_list_size 0 -hls_segment_type mpegts  -hls_flags single_file+round_durations \
    -f hls $output_path/video.m3u8

which ffmpeg

ffmpeg -y -vsync 0 -c:v hevc_cuvid -i $input_video -vf "scale=$resize_width_height,format=yuv420p"  \
    -map 0:v:0 -c:v h264_nvenc -b:v $video_bit_rate -g 30 -strict -2 \
    -color_primaries 1 -color_trc 1 -colorspace 1 \
    -hls_time $hls_time -hls_list_size 0 -hls_segment_type mpegts -hls_flags single_file+round_durations \
    -f hls $output_path/video.m3u8
