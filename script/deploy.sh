#!/bin/bash

az storage blob sync \
    --account-name lowplynet \
    --source ./_site \
    --container '$web/warmcolors' \
    --exclude-path assets \
    --delete-destination true
