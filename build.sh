#!/usr/bin/env bash

set -euf -o pipefail

go build github.com/skipio-corp/gr/cli/client
go build github.com/skipio-corp/gr/cli/server
