#!/bin/bash
set -e

# --
export PATH=$(go env GOPATH)/bin:$PATH
# --

source ./DevOps_Shell/read_config.sh

# --
# Call this function from './DevOps_Shell/read_config.yaml.sh' to get ES_HOST value in config.yaml file
get_value_from_yaml
# --

# Activate virtualenv && run serivce
SCRIPTDIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

# --
# swag init
# Finally, it’s time to generate the docs! All you need is one command —
swag init

# --
# Waitng for ES
./wait_for_es.sh $ES_HOST

if [[ -z "$PORT" ]]; then
    echo "Variable is empty"
    export PORT=9081
    echo $PORT
fi

# export ES_HOST=9206
# echo $ES_HOST

go run ./swagger.go
# ./swagger
