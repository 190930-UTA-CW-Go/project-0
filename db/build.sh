#!/bin/bash

# stop container
docker stop pgcontainer

# remove container
docker container rm pgcontainer

# remove image
docker rmi pgimage

# build image
docker build -t pgimage .

# create container
docker run --name=pgcontainer -p=5432:5432 -e POSTGRES_PASSWORD=password -d pgimage