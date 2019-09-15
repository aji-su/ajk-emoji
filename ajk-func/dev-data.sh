#!/bin/sh

until (mc config host add myminio http://minio:9000 $MINIO_ACCESS_KEY $MINIO_SECRET_KEY) do sleep 1; done
mc mb --region $MINIO_REGION myminio/test-bucket
mc policy set public myminio/test-bucket
tail -f /dev/null
