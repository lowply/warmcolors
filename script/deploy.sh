#!/bin/bash

aws \
    s3 sync \
    --exclude "assets/*" \
    --delete \
    ./_site/ s3://lowply.net/warmcolors/
