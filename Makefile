# thanx to https://makefiletutorial.com
#

# which tags (on features or scenarios) to execute
#TAGS = "@wip, not @skip"
#TAGS = "not @ignore"
TAGS = ""

# store location of test-results in variable
REPORTS ?= test-results/
FEATURES := $(shell find features -name '*.feature')


.PHONY: all clean features bdd report simple-report detailed-report bdd list json

features:
	echo $(FEATURES)

all: clean report


clean:
	rm -r $(REPORTS)*

#.PHONY: bdd
bdd:
	-godog --tags=$(TAGS)


report: simple-report detailed-report


simple-report: $(REPORTS)cucumber-report.json
	node ./assets/simple-cucumber-report.js


detailed-report: $(REPORTS)cucumber-report.json
	node ./assets/detailed-cucumber-report.js

json: $(REPORTS)cucumber-report.json

#$(REPORTS)cucumber-report.json: $(wildcard $(FEATURES)*.feature)
$(REPORTS)cucumber-report.json: $(FEATURES)
	-godog --format cucumber:$(REPORTS)cucumber-report.json --tags=$(TAGS)


# list command: https://stackoverflow.com/a/26339924/1782149

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
