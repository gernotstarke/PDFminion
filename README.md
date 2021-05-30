![](PDFminion-logo-small.png)
# PDFminion


Helper (_minion_) for some mundane tasks with PDF documents.


## Status
[![feature-linter](https://github.com/gernotstarke/PDFminion/actions/workflows/feature-linter.yml/badge.svg)](https://github.com/gernotstarke/PDFminion/actions/workflows/feature-linter.yml)
[![cucumber-report](https://github.com/gernotstarke/PDFminion/actions/workflows/cucumber-report.yml/badge.svg)](https://github.com/gernotstarke/PDFminion/actions/workflows/cucumber-report.yml)
[![go_test](https://github.com/gernotstarke/PDFminion/actions/workflows/go_test.yml/badge.svg)](https://github.com/gernotstarke/PDFminion/actions/workflows/go_test.yml)

[![Go Report Card](https://goreportcard.com/badge/github.com/gernotstarke/PDFminion)](https://goreportcard.com/report/github.com/gernotstarke/PDFminion)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gernotstarke_PDFminion&metric=alert_status)](https://sonarcloud.io/dashboard?id=gernotstarke_PDFminion)
[![Maintainability](https://api.codeclimate.com/v1/badges/c481ef8142826f71ff65/maintainability)](https://codeclimate.com/github/gernotstarke/PDFminion/maintainability)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Why PDFminion? 


## Development

We're using BDD (behavior driven development) with Cucumber to specify at least some of the requirements as _scenarios_.
These scenarios can be executed, similar to automated unit tests.

* [Godog](https://github.com/cucumber/godog), the official Cucumber tool
* [cucumber-html-reporter](https://www.npmjs.com/package/cucumber-html-reporter)


### Godog
 
````shell
go get github.com/cucumber/godog/cmd/godog@v0.11.0
````
### Cucumber HTML Reporter

It's written in JavaScript and requires `npm` and `node` to be available on your machine.

```shell
npm install cucumber-html-reporter --save-dev
```

## Usage of Development Tools

I squeezed the required commands into the file `test-and-create-report.sh`
(and into a github action )

### Godog

```shell
godog --format cucumber:test-results-results/cucumber-report.json
```

Notes: 

* godog requires features and scenarios to be written in a `features` directory.
* the `--format` switch can take a file and/or directory name


### Cucumber Report

```shell
node cucumber-report-index.js
```

Notes:

* cucumber-html-reporter can handle "Metadata" (see `cucumber-report-index.js`), for really interesting data
you should provide it with command line parameters, like commit-id
  
