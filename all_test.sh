#!/bin/bash

set -e

go clean -testcache

TEST_TARGETS=($(go list ./...))
for target in "${TEST_TARGETS[@]}"; do
    go test ${target}
done
