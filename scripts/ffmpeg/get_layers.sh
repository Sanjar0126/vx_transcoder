#!/bin/bash
ffprobe -v error -show_streams sample-avi-file.avi -print_format json
