#!/bin/bash
input_video=$1
destination=$2

aws s3 sync "$input_video" "$destination"

bot_token='1714608232:AAEI8xKe0dZ3Ouhx3GUCJbxJcJOumZGHg7I'
bot_chatID='-1001747193573'

input=$(basename $input_video)

message="%23transcoder%0AAws%20s3%20copy%20done:%0AInput%20File:%20${input}%0ADestination:%20${destination}"

send_text="https://api.telegram.org/bot${bot_token}/sendMessage?chat_id=${bot_chatID}&parse_mode=Markdown&text=${message}"

curl $send_text