#!/bin/bash
output_path=$1
audio_langs=$2
subtitle_langs=$3
resolutions=$4
video_measurements=$5

val_of_key() {
    case $1 in
        'ru') echo 'Russian - Русский';;
        'en') echo 'English - Английский';;
        'rus') echo 'Russian - Русский';;
        'eng') echo 'English - Английский';;
        'ukr') echo 'Ukraine - Украинский';;
        'track_1') echo 'Uzb - Узбекский';;
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
        '240p') echo '358400';;
        '360p') echo '512000';;
        '480p') echo '819200';;
        '720p') echo '1572864';;
        '1080p') echo '3145728';;
        '4K') echo '16777216';;
        *) echo 'Default setting - Стандартная настройка';;
    esac
}

generate_audio(){
    local input=$audio_langs

    # Set comma as the delimiter
    IFS=','

    #Read the split words into an array based on comma delimiter
    read -a splitted_input <<< $input

    for val in "${splitted_input[@]}";
    do
        if [ $val == $splitted_input ]
        then
            echo \#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID=\""stereo"\",LANGUAGE=\""$val"\",NAME=\""$(val_of_key $val)"\",DEFAULT=YES,AUTOSELECT=YES,URI=\""audios/$val/audio.m3u8"\" >> $output_path/master.m3u8
        else
            echo \#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID=\""stereo"\",LANGUAGE=\""$val"\",NAME=\""$(val_of_key $val)"\",DEFAULT=NO,AUTOSELECT=YES,URI=\""audios/$val/audio.m3u8"\" >> $output_path/master.m3u8
        fi
    done

    echo \ >> $output_path/master.m3u8
}

generate_subtitles(){
    local input=$subtitle_langs

    # Set comma as the delimiter
    IFS=','

    #Read the split words into an array based on comma delimiter
    read -a splitted_input <<< $input

    for val in "${splitted_input[@]}";
    do
        if [ $val == $splitted_input ]
        then
            echo \#EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID=\""subs"\",NAME=\""$(val_of_key $val)"\",DEFAULT=YES,AUTOSELECT=YES,FORCED=NO,LANGUAGE=\""$val"\",URI=\""subtitles/$val/sub.m3u8"\" >> $output_path/master.m3u8
        else
            echo \#EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID=\""subs"\",NAME=\""$(val_of_key $val)"\",DEFAULT=NO,AUTOSELECT=YES,FORCED=NO,LANGUAGE=\""$val"\",URI=\""subtitles/$val/sub.m3u8"\" >> $output_path/master.m3u8
        fi
    done

    echo \ >> $output_path/master.m3u8
}

generate_video(){
    local input=$resolutions
    local resolutions=$video_measurements

    # Set comma as the delimiter
    IFS=','

    #Read the split words into an array based on comma delimiter
    read -a splitted_input <<< $input
    read -a splitted_resolutions <<< $resolutions

    arraylength=${#splitted_input[@]}

    for (( i=0; i<${arraylength}; i++ ));
    do
        if [ -z "$subtitle_langs"]; then
            echo $i
            echo \#EXT-X-STREAM-INF:BANDWIDTH=$(val_of_key ${splitted_input[$i]}),CODECS=\""avc1.42c00d,mp4a.40.2\"",RESOLUTION=${splitted_resolutions[$i]},AUDIO=\""stereo"\" >> $output_path/master.m3u8
            echo videos/${splitted_input[$i]}/video.m3u8 >> $output_path/master.m3u8
        else
            echo $i
            echo \#EXT-X-STREAM-INF:BANDWIDTH=$(val_of_key ${splitted_input[$i]}),CODECS=\""avc1.42c00d,mp4a.40.2\"",RESOLUTION=${splitted_resolutions[$i]},AUDIO=\""stereo"\",SUBTITLES=\""subs"\" >> $output_path/master.m3u8
            echo videos/${splitted_input[$i]}/video.m3u8 >> $output_path/master.m3u8
        fi
    done
}

init_master_playlist(){
    echo "#EXTM3U
#EXT-X-VERSION:3" > $output_path/master.m3u8
}

init_master_playlist
generate_audio
generate_subtitles
generate_video
