# !/bin/bash

docker build . -t indexer
docker run --rm indexer
