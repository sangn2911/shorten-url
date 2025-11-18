#!/bin/bash
rm alwaysdata/shortlink
docker build -f alwaysdata/Dockerfile.binary -t shortlink:latest .
container_id=$(docker create shortlink:latest)
docker cp $container_id:/app/shortlink ./alwaysdata/shortlink
docker rm $container_id
docker rmi shortlink:latest