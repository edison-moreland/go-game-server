#!/usr/bin/env bash
# This exists to start locust from GoLand IDE

# $1 = path to venv
source ${1}/bin/activate

# $2 = host to test
locust --host=${2}