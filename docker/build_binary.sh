#!/bin/bash
docker build -f docker/Dockerfile.binary -t shortlink:latest .
container_id=$(docker create shortlink:latest)
docker cp $container_id:/app/shortlink ./shortlink
docker rm $container_id
docker rmi shortlink:latest