#!/bin/bash
input_video=$1
destination=$2

aws s3 sync "$input_video" "$destination"
