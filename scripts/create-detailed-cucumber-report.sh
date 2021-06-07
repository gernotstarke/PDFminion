#!/bin/zsh

echo

REPORT_PATH="test-results/cucumber-report"

# clear old report
rm ${REPORT_PATH}.json


# run godog to create new bdd report
#godog --format cucumber:${REPORT_PATH}.json


# create HTML report
node ./assets/detailed-cucumber-report.js

