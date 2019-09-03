#!/bin/sh

export AWS_ACCESS_KEY_ID=testkey
export AWS_SECRET_ACCESS_KEY=testsecret
export AWS_DEFAULT_REGION=ap-northeast-1

docker-compose up -d

aws s3api create-bucket \
  --bucket=test-bucket \
  --endpoint-url=http://localhost:9000
