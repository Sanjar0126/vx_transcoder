#!/bin/bash
input_video=$1
lang=$2
audio_bit_rate=$3
audio_track_number=$4
segment_time=$5
output_path=$6

echo -y -i $input_video -b:a $audio_bit_rate -map 0:a:$audio_track_number \
    -ac 2 -c:a aac -strict 2 -f segment -segment_time $segment_time \
    -segment_list_type m3u8 -segment_list \
    $output_path/audios/$lang/audio.m3u8 \
    $output_path/audios/$lang/audio_%d.ts

ffmpeg \
    -y -i $input_video -b:a $audio_bit_rate -map 0:a:$audio_track_number \
    -ac 2 -c:a aac -strict 2 -threads 12 -f segment -segment_time $segment_time \
    -segment_list_type m3u8 -segment_list \
    $output_path/audios/$lang/audio.m3u8 \
    $output_path/audios/$lang/audio_%d.ts
