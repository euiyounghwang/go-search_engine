#!/bin/bash
set -e
source ./DevOps_Shell/read_config.sh

# --
# Call this function from './DevOps_Shell/read_config.yaml.sh' to get ES_HOST value in config.yaml file
get_value_from_yaml
# --

# Activate virtualenv && run serivce
SCRIPTDIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

# --
# swag init
swag init

# --
# Waitng for ES
./wait_for_es.sh $ES_HOST

go run ./swagger.go
# ./swagger
