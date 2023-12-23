#!/bin/bash

set -eu

#docker build --no-cache \


docker build \
  -f "$(dirname "$0")/Dockerfile" \
  -t go-elasticsearch-api:es \
  --target runtime \
  "$(dirname "$0")/."
