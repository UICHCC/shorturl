#!/bin/bash

docker stop shorturl
docker rm shorturl
docker rmi dengrenjie31/shorturl
docker compose up -d --remove-orphans