#!/bin/bash

find . -name .DS_Store -delete -exec echo removed: {} \;

aws \
    --profile s3backup \
    s3 sync \
    --exclude "main.css" \
    --exclude "main.css.map" \
    --delete \
    ./assets/ s3://lowply.org/warmcolors/assets/
