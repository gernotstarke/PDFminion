#!/bin/zsh

echo

REPORT_PATH="test-results/cucumber-report"

godog --format cucumber:${REPORT_PATH}.json


# create HTML report
node ./assets/detailed-cucumber-report.js

