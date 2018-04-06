#!/bin/bash

# A simple integration test which undertakes the whole lifecycle for creating
# a tmplinux environment using container

set -e

BAZEL_RUN="bazel run //:tmplinux"

integration_test::setup() {
  make clean
  make build
}

integration_test::run() {
  $BAZEL_RUN container validate
  $BAZEL_RUN container start

  $BAZEL_RUN container ssh &
  local pid=$!
  kill -9 $pid

  $BAZEL_RUN container stop
  $BAZEL_RUN container rm
}

integration_test::teardown() {
  local trap_code="$?"
  make clean
  exit $trap_code
}

trap 'integration_test::teardown' EXIT

integration_test::setup
integration_test::run
