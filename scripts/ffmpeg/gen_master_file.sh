#!/bin/bash

val_of_key() {
    case $1 in
        'ru') echo 'Russian - Русский';;
        'en') echo 'English - Английский';;
        'rus') echo 'Russian - Русский';;
        'eng') echo 'English - Английский';;
        'ukr') echo 'Ukraine - Украинский';;
        'track_1') echo 'Трек 1 - Track 1';;
        'track_2') echo 'Трек 2 - Track 2';;
        'track_3') echo 'Трек 3 - Track 3';;
        'track_4') echo 'Трек 4 - Track 4';;
        'track_5') echo 'Трек 5 - Track 5';;
        'track_6') echo 'Трек 6 - Track 6';;
        'rus-1') echo 'Russian - Русский 2';;
        'ukr-1') echo 'Ukraine - Украинский 2';;
        'eng-1') echo 'English - Английский 2';;
        'rus-2') echo 'Russian - Русский 3';;
        'eng-2') echo 'English - Английский 3';;
        'ukr-3') echo 'Ukraine - Украинский 3';;
        'rus-3') echo 'Russian - Русский 4';;
        'eng-3') echo 'English - Английский 4';;
        'rus-4') echo 'Russian - Русский 5';;
        'ukr-4') echo 'Ukraine - Украинский 4';;
        'eng-4') echo 'English - Английский 5';;
        'rus-5') echo 'Russian - Русский 6';;
        'ukr-5') echo 'Ukraine - Украинский 6';;
        'eng-5') echo 'English - Английский 6';;
        'rus-6') echo 'Russian - Русский 7';;
        'ukr-6') echo 'Ukraine - Украинский 7';;
        'eng-6') echo 'English - Английский 7';;
        'rus-7') echo 'Russian - Русский 8';;
        'ukr-7') echo 'Ukraine - Украинский 8';;
        'eng-7') echo 'English - Английский 8';;
        'rus-8') echo 'Russian - Русский 9';;
        'ukr-8') echo 'Ukraine - Украинский 9';;
        'eng-8') echo 'English - Английский 9';;
        'rus-9') echo 'Russian - Русский 10';;
        'ukr-9') echo 'Ukraine - Украинский 10';;
        'eng-9') echo 'English - Английский 10';;
        *) echo 'Default setting - Стандартная настройка';;
    esac
}

escape() {
  printf "\n" >> "$output_path"/master.m3u8
}

get_default() {
  case $1 in
      0) echo 'YES';;
      *) echo 'NO';;
    esac
}

create_audio() {
  escape

  local audio_langs=$audio_langs

  IFS=, read -ra audio_array <<< "$audio_langs"

  local count=0
  for audio in "${audio_array[@]}"
  do
    printf "#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID=\"%s\",LANGUAGE=\"%s\",NAME=\"%s\",DEFAULT=%s,AUTOSELECT=YES,URI=\"audios/%s/audio.m3u8\"\n" "$audio_group" "$audio" "$(val_of_key $audio)" "$(get_default "$count")" "$audio" >> "$output_path"/master.m3u8
    ((count++))
  done
}

create_sub() {
  escape

  local subs="$subtitle_langs"

  IFS=, read -ra sub_array <<< "$subs"

  local count=0
  for sub in "${sub_array[@]}"
  do
    printf "#EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID=\"%s\",LANGUAGE=\"%s\",NAME=\"%s\",DEFAULT=%s,AUTOSELECT=YES,FORCED=NO,URI=\"subtitles/%s/sub.m3u8\"\n" "$subtitle_group" "$sub" "$(val_of_key "$sub")" "$(get_default "$count")" "$sub" >> "$output_path"/master.m3u8
    ((count++))
  done
}

audio_group="audio"
subtitle_group="subtitle"
output_path=$1
audio_langs=$2
subtitle_langs=$3
resolutions=$4
video_measurements=$5

rm -f "$output_path"/master.m3u8
printf "#EXTM3U\n#EXT-X-VERSION:3\n" > "$output_path"/master.m3u8
create_audio "$audio_langs"
create_sub "$subtitle_langs"
