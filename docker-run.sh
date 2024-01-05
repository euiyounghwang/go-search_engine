#!/bin/bash

set -eu

SCRIPTDIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
echo $SCRIPTDIR


docker run --rm -it -d \
  --name go-search_engine-api --publish 9077:9081 --expose 9081 \
  --network bridge \
  -e ES_HOST=http://host.docker.internal:9209 \
  -v "$SCRIPTDIR:/app" \
  go-search_engine-api:es


# docker run --rm -it -d \
#   --name go-search_engine-api --publish 9088:9080 --expose 9080 \
#   --network bridge \
#   -e ES_HOST=http://host.docker.internal:9209 \
#   -v "$SCRIPTDIR:/app" \
#   go-search_engine-api:es