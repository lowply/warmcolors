#!/bin/bash

aws --profile private s3api put-bucket-website --bucket warmcolors.info --website-configuration file://bucket_website_hosting.json
aws --profile private s3api put-bucket-policy --bucket warmcolors.info --policy file://bucket_policy.json
