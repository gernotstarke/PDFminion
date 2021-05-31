#!/bin/zsh

echo on


COMMIT_ID=$(git rev-parse --verify HEAD)

godog --format cucumber:test-results/cucumber-report.json

# create HTML report
node ./cucumber-report-index.js


# publish report to cucumber-studio

# for cucumber-studio we need to convert json to messages:
#cat test-results/cucumber-report.json | ~/node_modules/.bin/json-to-messages > test-results/cucumber-report.messages

# now use curl to publish
#curl -X POST \
#      https://studio.cucumber.io/cucumber_project/results \
#      -F messages=@<path to your result file> \
#      -H "project-access-token: 27877807453784791111842877191917738921804894910830921354" \
#      -H "provider: github" \
#      -H "repo: gernotstarke/golang-bdd" \
#      -H "branch: main" \
#      -H "revision: ${COMMIT_ID}"
