#!/bin/bash

# Simple script to switch wallpapers using swww
# Usage: ./wallpaper.sh /path/to/image_or_video

if [ -z "$1" ]; then
    echo "Usage: $0 /path/to/image_or_video"
    exit 1
fi

swww img "$1" --transition-type random --transition-step 90 --transition-fps 60
