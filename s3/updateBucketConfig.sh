#!/bin/bash

aws --profile private s3api put-bucket-policy --bucket warmcolors --policy file://bucket_policy.json
