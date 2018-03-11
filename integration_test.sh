#!/bin/bash

# A simple integration test which undertakes the whole lifecycle for creating
# a tmplinux environment using container

set -e

integration_test::setup() {
  make build
}

integration_test::run() {
  ./tmplinux container validate
  ./tmplinux container start

  ./tmplinux container ssh &
  local pid=$!
  kill -9 $pid

  ./tmplinux container stop
  ./tmplinux container rm
}

integration_test::teardown() {
  make clean
}

trap 'integration_test::teardown' EXIT

integration_test::setup
integration_test::run
