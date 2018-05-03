#/bin/bash

find . -name .DS_Store -delete -exec echo removed: {} \;
aws --profile private s3 --no-follow-symlinks sync ./assets/ s3://warmcolors/radio/
