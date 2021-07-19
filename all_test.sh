#!/bin/bash

set -e

go clean -testcache

text='
aaa
bbb
ccc
'
array_text=(`echo $text`)

TEST_TARGETS=($(go list ./...))
for target in "${TEST_TARGETS[@]}"; do
    go test ${target}
done
