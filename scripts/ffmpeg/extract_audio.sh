#!/bin/bash
input_video=$1
audio_bit_rate=$2
audio_track_number=$3
segment_time=$4
output_path=$5

echo -y -i $input_video -b:a $audio_bit_rate -map 0:a:$audio_track_number \
    -ac 2 -c:a aac -strict 2 -f segment -segment_time $segment_time \
    -segment_list_type m3u8 -segment_list \
    $output_path/audio.m3u8 \
    $output_path/audio_%d.ts

ffmpeg \
    -y -i $input_video -b:a $audio_bit_rate -map 0:a:$audio_track_number \
    -ac 2 -c:a aac -strict 2 -threads 12 -f segment -segment_time $segment_time \
    -segment_list_type m3u8 -segment_list \
    $output_path/audio.m3u8 \
    $output_path/audio_%d.ts
