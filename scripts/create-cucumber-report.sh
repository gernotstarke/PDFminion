#!/bin/zsh

echo on

REPORT_PATH="test-results/simple-cucumber-report"

godog --format cucumber:${REPORT_PATH}.json


# create HTML report
node ./assets/simple-cucumber-report.js

