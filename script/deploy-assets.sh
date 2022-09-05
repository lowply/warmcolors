#!/bin/bash

find . -name .DS_Store -delete -exec echo removed: {} \;

az storage blob sync \
    --account-name lowplynet \
    --source ./assets \
    --container '$web/warmcolors/assets' \
    --delete-destination true
