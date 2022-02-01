#!/bin/bash

aws \
    s3 sync \
    --exclude "assets/img/*" \
    --exclude "assets/radio/*" \
    --delete \
    ./_site/ s3://lowply.org/warmcolors/
