# ![PDFminion](PDFminion-logo.png)


Helper (_minion_) for some mundane tasks with PDF documents, among others:

* add page numbers
* add header and/or footer text
* concatenate (combine multiple PDF files into a single file)

It shall have a (multi-platform) graphical user interface, at least for Mac-OS, Windows and maybe Linux.

>  minion: a servile dependent, follower, or underling.<br> 
> "He's one of the boss's minions."<br>
> From: [Merriam-Webster Dictionary](https://www.merriam-webster.com/dictionary/minion)

## Status
[![feature-linter](https://github.com/gernotstarke/PDFminion/actions/workflows/feature-linter.yml/badge.svg)](https://github.com/gernotstarke/PDFminion/actions/workflows/feature-linter.yml)
[![go_test](https://github.com/gernotstarke/PDFminion/actions/workflows/go_test.yml/badge.svg)](https://github.com/gernotstarke/PDFminion/actions/workflows/go_test.yml)

[![Go Report Card](https://goreportcard.com/badge/github.com/gernotstarke/PDFminion)](https://goreportcard.com/report/github.com/gernotstarke/PDFminion)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gernotstarke_PDFminion&metric=alert_status)](https://sonarcloud.io/dashboard?id=gernotstarke_PDFminion)
[![Maintainability](https://api.codeclimate.com/v1/badges/c481ef8142826f71ff65/maintainability)](https://codeclimate.com/github/gernotstarke/PDFminion/maintainability)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Why PDFminion? 


## Development

We're using BDD (behavior driven development) with Cucumber to specify at least part of the requirements as _scenarios_.
These scenarios can be executed, similar to automated unit tests.

* [Godog](https://github.com/cucumber/godog), the official Cucumber tool
* [Cucumber HTML Reporter](https://www.npmjs.com/package/cucumber-html-reporter)
* [Cucumber Multi Reporter (more detailed)](https://github.com/wswebcreation/multiple-cucumber-html-reporter)

Use `./create-detailed-cucumber-report.sh` to generate a detailed BDD Cucumber report.

### Deviation from Standard golang practices
As of June 2021, the `godog` bdd tool does not respect the standard golang layout practice
of putting test files next to the tested-code.
Instead, the step definitions need to be present in the root folder!

To avoid confusion, I prefixed the step definitions with `stepdef_` - so they are easily recognizable.
Other (non-bdd/cucumber) automated tests will reside within the appropriate package folders.

See this [Cucumber/godog issue](https://github.com/cucumber/godog/issues/373).

### Godog
 
````shell
go get github.com/cucumber/godog/cmd/godog@v0.12.0
````
### Cucumber HTML Reporter

It's written in JavaScript and requires `npm` and `node` to be available on your machine.

```shell
npm install cucumber-html-reporter --save-dev
```

### Cucumber Multi Reporter

Again, JavaScript, see above:

```shell
npm install multiple-cucumber-html-reporter --save-dev
```


## Usage of Development Tools

I squeezed the required commands into the files `create-cucumber-report.sh`
and `create-detailed-cucumber-report.sh`.

### Godog

```shell
godog --format cucumber:test-results-results/cucumber-report.json
```

Notes: 

* godog requires features and scenarios to be written in a `features` directory.
* the `--format` switch can take a file and/or directory name


### Cucumber Report

```shell
node ./assets/simple-cucumber-report.js
```

