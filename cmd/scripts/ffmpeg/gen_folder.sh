#!/bin/bash

create_folder() {
  local files=$2
  local dest=$1
  echo "$files"
  IFS=, read -ra file_array <<< "$files"

  for file in "${file_array[@]}"
  do
    mkdir -p "$dest/$file"
  done
}

output=$1
audio_list=$2
video_list=$3
subtitle_list=$4

echo "$output"
echo "$audio_list"
echo "$video_list"
echo "$subtitle_list"

mkdir -p "$output"

create_folder  "$output/audios" "$audio_list"
create_folder  "$output/videos" "$video_list"
create_folder  "$output/subtitles" "$subtitle_list"


