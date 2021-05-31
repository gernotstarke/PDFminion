#!/bin/zsh

echo on

REPORT_PATH="test-results/cucumber-report"

godog --format cucumber:${REPORT_PATH}.json


# create HTML report
node ./assets/cucumber-report-index.js

