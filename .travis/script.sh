#!/usr/bin/env bash

set -e

go get -t ./...
if [ ${TESTMODE} == "unit" ]; then
  ginkgo -r --cover --randomizeAllSpecs --randomizeSuites --trace --progress --skipPackage integrationtests --skipMeasurements
fi

if [ ${TESTMODE} == "integration" ]; then
  # run benchmark tests
  ginkgo --randomizeAllSpecs --randomizeSuites --trace --progress -focus "Benchmark"
  ginkgo --race --randomizeAllSpecs --randomizeSuites --trace --progress -focus "Benchmark"
  # run integration tests
  ginkgo -r --randomizeAllSpecs --randomizeSuites --trace --progress integrationtests
fi
