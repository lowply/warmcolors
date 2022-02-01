#!/bin/bash

find . -name .DS_Store -delete -exec echo removed: {} \;

aws \
    --profile s3backup \
    s3 sync \
    --exclude "assets/img/*" \
    --exclude "assets/radio/*" \
    --delete \
    ./_site/ s3://lowply.org/warmcolors/