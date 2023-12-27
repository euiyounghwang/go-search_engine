#!/bin/bash

SCRIPTDIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

# cd /Users/euiyoung.hwang/go/pkg/mod/github.com/smartystreets/goconvey@v1.8.1
# cd /Users/euiyoung.hwang/go/bin

# /Users/euiyoung.hwang/go/bin/goconvey --help
/Users/euiyoung.hwang/go/bin/goconvey --workDir=$SCRIPTDIR/tests --port=7090