#!/bin/bash

set -e

# Activate virtualenv && run serivce

# cd /Users/euiyoung.hwang/ES/Python_Workspace/python-django
SCRIPTDIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

cd $SCRIPTDIR/tests

go test -v