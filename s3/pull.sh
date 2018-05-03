#/bin/bash

aws --profile private s3 --no-follow-symlinks --delete sync s3://warmcolors/radio/ ./assets/ 
