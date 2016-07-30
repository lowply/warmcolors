#!/bin/bash

aws --profile private s3api put-bucket-website --bucket warmcolors.info --website-configuration file://bucket.json
