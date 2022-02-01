#!/bin/bash

find . -name .DS_Store -delete -exec echo removed: {} \;

aws \
    --profile s3backup \
    s3 sync \
    --delete \
    ./assets/ s3://lowply.net/warmcolors/assets/
